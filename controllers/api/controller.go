package api

import (
	"fmt"
	"github.com/astaxie/beego"
	"strings"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/util"
)

type ApiController struct {
	beego.Controller
}

func (c ApiController) GetAuth() (*security.UserClaims, error) {
	if claims, err := security.ParseAuthHeader(c.Controller); err != nil {
		return nil, err
	} else {
		return claims, err
	}
}

func (c ApiController) SerializeData(data interface{}) {
	panic("serialize function not define")
}

func (c ApiController) GetPage() (page int64, pageSize int64) {
	page, pageSize = util.ParsePageRequest(c.Controller)
	return
}

func (c ApiController) CheckPermission(permissions []ApiPermissionInterface, context map[string]interface{}) error {
	for _, permission := range permissions {
		if hasPermission := permission.CheckPermission(context); !hasPermission {
			return PermissionDeniedError
		}
	}
	return nil
}

func (c ApiController) ServerPageResult(result []interface{}, count int64,Page int64,PageSize int64) {
	response := util.PageResponse{
		Count:    count,
		Page:     Page,
		PageSize: PageSize,
		Result:result,
	}
	if count > Page*PageSize {
		urlString := fmt.Sprintf("%s%s", util.GetSiteAndPortUrl(c.Controller), c.Ctx.Input.URI())
		nextPage := strings.Replace(urlString, fmt.Sprint("page=", Page), fmt.Sprint("page=", Page+1), 1)
		response.NextPage = &nextPage
	}
	if Page > 1 {
		urlString := fmt.Sprintf("%s%s", util.GetSiteAndPortUrl(c.Controller), c.Ctx.Input.URI())
		prevPage := strings.Replace(urlString, fmt.Sprint("page=", Page), fmt.Sprint("page=", Page-1), 1)
		response.PrevPage = &prevPage
	}
	c.Data["json"] = response
	c.ServeJSON()
}
