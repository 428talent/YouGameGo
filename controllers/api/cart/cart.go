package cart

import (
	"reflect"
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
	results := make([]interface{}, 0)
	for _, item := range cartItems {
		results = append(results, reflect.ValueOf(*item).Interface())
	}

	serializerDataList := serializer.SerializeMultipleData(&serializer.CartModel{}, results, util.GetSiteAndPortUrl(c.Controller))
	c.ServerPageResult(serializerDataList, count, page, pageSize)

}
