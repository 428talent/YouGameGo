package api

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/copier"
	"reflect"
	"strconv"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
	"yougame.com/yougame-server/util"
)

type View interface {
	Exec() error
}

type ListView struct {
	Controller    *ApiController
	Init          func()
	QueryBuilder  service.ApiQueryBuilder
	ModelTemplate serializer.Template
	GetTemplate   func() serializer.Template
	GetPage       func() (page int64, pageSize int64)
	SetFilter     func(builder service.ApiQueryBuilder)
	OnGetResult   func(interface{})
}

func (v *ListView) Exec() error {
	if v.Init != nil {
		v.Init()
	}
	page, pageSize := v.Controller.GetPage()

	serializeTemplate := v.ModelTemplate
	if v.GetTemplate != nil {
		serializeTemplate = v.GetTemplate()
	}
	if v.GetPage != nil {
		v.QueryBuilder.SetPage(v.GetPage())
	} else {
		v.QueryBuilder.SetPage(page, pageSize)
	}
	if v.SetFilter != nil {
		v.SetFilter(v.QueryBuilder)
	}
	count, modelList, err := v.QueryBuilder.ApiQuery()
	if err != nil {
		return err
	}
	if v.OnGetResult != nil {
		v.OnGetResult(modelList)
	}
	result := serializer.SerializeMultipleTemplate(modelList, serializeTemplate, map[string]interface{}{
		"site": util.GetSiteAndPortUrl(v.Controller.Controller),
	})
	v.Controller.ServerPageResult(result, *count, page, pageSize)
	return nil
}

type ObjectView struct {
	Controller          *ApiController
	QueryBuilder        service.ApiQueryBuilder
	ModelTemplate       serializer.Template
	LookUpField         string
	Model               interface{}
	SerializeContext    map[string]interface{}
	GetTemplate         func() serializer.Template
	SetFilter           func(builder service.ApiQueryBuilder)
	OnGetResult         func(model interface{})
	SetSerializeContext func(context map[string]interface{})
}

func (v *ObjectView) Exec() error {
	if v.LookUpField != "-" {
		lookup := ":id"
		if len(v.LookUpField) > 0 {
			lookup = v.LookUpField
		}
		lookUpParam := v.Controller.Ctx.Input.Param(lookup)
		id, err := strconv.Atoi(lookUpParam)
		if err != nil {
			return err
		}
		v.QueryBuilder.InId(id)
	}

	if v.SetFilter != nil {
		v.SetFilter(v.QueryBuilder)
	}

	count, resultSet, err := v.QueryBuilder.ApiQuery()
	if err != nil {
		return err
	}
	if *count == 0 {
		return ResourceNotFoundError
	}

	v.SerializeContext = map[string]interface{}{
		"site": util.GetSiteAndPortUrl(v.Controller.Controller),
	}
	if v.SetSerializeContext != nil {
		v.SetSerializeContext(v.SerializeContext)
	}
	if v.ModelTemplate == nil && v.GetTemplate == nil {
		return errors.New("not set serialize template")
	}
	if v.GetTemplate != nil {
		v.ModelTemplate = v.GetTemplate()
	}
	if v.OnGetResult != nil {
		v.OnGetResult(reflect.ValueOf(resultSet).Index(0).Interface())
	}
	v.ModelTemplate.Serialize(reflect.ValueOf(resultSet).Index(0).Interface(), v.SerializeContext)
	v.Controller.Data["json"] = v.ModelTemplate
	v.Controller.ServeJSON()
	return nil
}

type DeleteView struct {
	Controller           *ApiController
	Init                 func()
	Model                models.DataModel
	Permissions          []PermissionInterface
	GetPermissionContext func(permissionContext *map[string]interface{}) *map[string]interface{}
}

func (v *DeleteView) Exec() error {
	claims, err := v.Controller.GetAuth()
	if err != nil {
		return ClaimsNoFoundError
	}
	if claims == nil {
		return ClaimsNoFoundError
	}
	permissionContext := map[string]interface{}{
		"claims": claims,
	}
	if v.GetPermissionContext != nil {
		v.GetPermissionContext(&permissionContext)
	}
	err = v.Controller.CheckPermission(v.Permissions, permissionContext)
	if err != nil {
		return PermissionDeniedError
	}
	idParam := v.Controller.Ctx.Input.Param(":id")
	if len(idParam) == 0 {
		return ResourceNotFoundError
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		panic(err)
	}

	err = v.Model.Query(int64(id))
	if err != nil {
		return err
	}

	err = service.DeleteData(v.Model)
	if err != nil {
		panic(err)
	}

	v.Controller.ResponseWithSuccess()
	return nil
}

type UpdateView struct {
	Controller           *ApiController
	Init                 func()
	Parser               interface{}
	Model                models.DataModel
	Permissions          []PermissionInterface
	ModelTemplate        serializer.Template
	GetTemplate          func() serializer.Template
	GetPermissionContext func(permissionContext *map[string]interface{}) *map[string]interface{}
}

func (v *UpdateView) Exec() error {
	claims, err := v.Controller.GetAuth()
	if err != nil {
		return ClaimsNoFoundError
	}
	if claims == nil {
		return ClaimsNoFoundError
	}
	permissionContext := map[string]interface{}{
		"claims": claims,
	}
	if v.GetPermissionContext != nil {
		v.GetPermissionContext(&permissionContext)
	}
	err = v.Controller.CheckPermission(v.Permissions, permissionContext)
	if err != nil {
		return PermissionDeniedError
	}

	err = json.Unmarshal(v.Controller.Ctx.Input.RequestBody, v.Parser)
	if err != nil {
		return ParseJsonDataError
	}

	jsonObjectMap := map[string]interface{}{}
	err = json.Unmarshal([]byte(v.Controller.Ctx.Input.RequestBody), &jsonObjectMap)
	if err != nil {
		return ParseJsonDataError
	}
	modelId, err := strconv.Atoi(v.Controller.Ctx.Input.Param(":id"))
	if err != nil {
		return err
	}

	updateFields := make([]string, 0)
	for field := range jsonObjectMap {
		updateFields = append(updateFields, field)
	}

	err = copier.Copy(v.Model, v.Parser)
	if err != nil {
		return err
	}
	err = service.UpdateData(int64(modelId), v.Model, updateFields...)
	if err != nil {
		return err
	}

	err = v.Model.Query(int64(modelId))
	if err != nil {
		return err
	}

	if v.GetTemplate != nil {
		v.ModelTemplate = v.GetTemplate()
	}

	v.ModelTemplate.Serialize(v.Model, map[string]interface{}{
		"site": util.GetSiteAndPortUrl(v.Controller.Controller),
	})

	v.Controller.Data["json"] = v.ModelTemplate
	v.Controller.ServeJSON()
	return nil
}

type CreateView struct {
	Controller           *ApiController
	Parser               interface{}
	Model                models.DataModel
	Permissions          []PermissionInterface
	ModelTemplate        serializer.Template
	GetTemplate          func() serializer.Template
	GetPermissionContext func(permissionContext *map[string]interface{}) *map[string]interface{}
	OnPrepareSave func(c *CreateView)
	Validate      func(v *CreateView)
	OnSave        func(v *CreateView) error
}

func (v *CreateView) Exec() error {
	claims, err := v.Controller.GetAuth()
	if err != nil {
		return ClaimsNoFoundError
	}
	if claims == nil {
		return ClaimsNoFoundError
	}
	permissionContext := map[string]interface{}{
		"claims": claims,
	}
	if v.GetPermissionContext != nil {
		v.GetPermissionContext(&permissionContext)
	}
	err = v.Controller.CheckPermission(v.Permissions, permissionContext)
	if err != nil {
		return PermissionDeniedError
	}

	err = json.Unmarshal(v.Controller.Ctx.Input.RequestBody, v.Parser)
	if err != nil {
		return ParseJsonDataError
	}

	if v.Validate != nil {
		v.Validate(v)
	}

	err = copier.Copy(v.Model, v.Parser)
	if err != nil {
		return err
	}

	if v.OnPrepareSave != nil {
		v.OnPrepareSave(v)
	}

	if v.OnSave != nil {
		err := v.OnSave(v)
		if err != nil {
			panic(err)
		}
	} else {
		err = service.SaveData(v.Model)
		if err != nil {
			panic(err)
		}
	}

	if v.GetTemplate != nil {
		v.ModelTemplate = v.GetTemplate()
	}

	v.ModelTemplate.Serialize(v.Model, map[string]interface{}{
		"site": util.GetSiteAndPortUrl(v.Controller.Controller),
	})

	v.Controller.Data["json"] = v.ModelTemplate
	v.Controller.ServeJSON()
	return nil

}
