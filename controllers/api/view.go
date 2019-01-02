package api

import (
	"errors"
	"reflect"
	"strconv"
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
	//beego.Debug(reflect.ValueOf(resultSet).Index(0).Interface())
	//data := reflect.ValueOf(resultSet).Index(0).Elem().Interface()
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
