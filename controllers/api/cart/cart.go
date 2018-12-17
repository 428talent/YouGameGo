package cart

import (
	"github.com/astaxie/beego"
	"strconv"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
	"yougame.com/yougame-server/util"
)

type ApiCartController struct {
	api.ApiController
}

func (c ApiCartController) GetCartList() {
	var err error
	defer api.CheckError(func(e error) {
		beego.Debug(e)
		api.HandleApiError(c.Controller, e)
	})
	claims, err := c.GetAuth()
	if err != nil {
		panic(security.ReadAuthorizationFailed)
	}
	if claims == nil {
		panic(security.ReadAuthorizationFailed)
	}

	cartUserId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		panic(err)
	}

	permissionContext := map[string]interface{}{
		"claims":     *claims,
		"cartUserId": cartUserId,
	}
	permissions := []api.ApiPermissionInterface{
		GetOtherCartPermission{},
		GetSelfCartPermission{},
	}
	err = c.CheckPermission(permissions, permissionContext)
	if err != nil {
		panic(err)
	}
	queryBuilder := service.CartQueryBuilder{}
	page, pageSize := c.GetPage()
	queryBuilder.SetPage(page,pageSize)
	user, err := models.GetUserById(claims.UserId)
	if err != nil {
		panic(security.ReadAuthorizationFailed)
	}
	queryBuilder.InUser(user.Id)
	count,cartItems,err := queryBuilder.Query()
	if err != nil {
		panic(err)
	}
	results:= serializer.SerializeMultipleTemplate(cartItems,&serializer.CartTemplate{},map[string]interface{}{
		"site":util.GetSiteAndPortUrl(c.Controller),
	})

	c.ServerPageResult(results, count, page, pageSize)

}
