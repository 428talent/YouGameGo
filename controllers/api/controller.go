package api

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
	"yougame.com/yougame-server/util"
)

type ApiController struct {
	beego.Controller
	User *models.User
	Role string
}

func (c *ApiController) GetAuth() (*security.UserClaims, error) {
	claims, err := security.ParseAuthHeader(c.Controller)
	if err != nil {
		return nil, err
	}

	user := service.GetUserById(claims.UserId)
	c.User = user
	return claims, nil
}

func (c ApiController) SerializeData(data interface{}) {
	panic("serialize function not define")
}

func (c ApiController) GetPage() (page int64, pageSize int64) {
	page, pageSize = util.ParsePageRequest(c.Controller)
	return
}

func (c ApiController) CheckPermission(permissions []PermissionInterface, context map[string]interface{}) error {
	for _, permission := range permissions {
		if hasPermission := permission.CheckPermission(context); !hasPermission {
			return PermissionDeniedError
		}
	}
	return nil
}

func (c ApiController) ServerPageResult(result interface{}, count int64, Page int64, PageSize int64) {
	response := util.PageResponse{
		Count:    count,
		Page:     Page,
		PageSize: PageSize,
		Result:   result,
	}
	urlString := fmt.Sprint(util.GetSiteAndPortUrl(c.Controller), c.Ctx.Input.URL())
	queryParams := c.Ctx.Request.URL.Query()
	if count > Page*PageSize {
		queryParams.Set("page", strconv.Itoa(int(Page+1)))
		nextPage := fmt.Sprint(urlString, "?", queryParams.Encode())
		response.NextPage = &nextPage
	}
	if Page > 1 {
		queryParams.Set("page", strconv.Itoa(int(Page-1)))
		prevPage := fmt.Sprint(urlString, "?", queryParams.Encode())
		response.PrevPage = &prevPage
	}
	c.Data["json"] = response
	c.ServeJSON()
}

func (c ApiController) WithErrorContext(doSomething func()) {
	defer CheckError(func(e error) {
		beego.Debug(e)
		HandleApiError(c.Controller, e)
	})
	doSomething()
}

type ListView struct {
	Controller    *ApiController
	QueryBuilder  service.ApiQueryBuilder
	ModelTemplate serializer.Template
	GetTemplate   func() serializer.Template
	GetPage       func() (page int64, pageSize int64)
	SetFilter     func(builder service.ApiQueryBuilder)
}

func (v *ListView) GetModelTemplate() serializer.Template {
	return v.ModelTemplate
}

func (v *ListView) Exec() error {
	page, pageSize := v.Controller.GetPage()

	serializeTemplate := v.ModelTemplate
	if v.GetTemplate != nil {
		serializeTemplate = v.GetTemplate()
	}
	if v.GetPage != nil {
		v.QueryBuilder.SetPage(v.GetPage())
	}
	if v.SetFilter != nil {
		v.SetFilter(v.QueryBuilder)
	}
	count, modelList, err := v.QueryBuilder.ApiQuery()
	if err != nil {
		return err
	}
	result := serializer.SerializeMultipleTemplate(modelList, serializeTemplate, map[string]interface{}{
		"site": util.GetSiteAndPortUrl(v.Controller.Controller),
	})
	v.Controller.ServerPageResult(result, *count, page, pageSize)
	return nil
}
