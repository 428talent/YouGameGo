package wishlist

import (
	"encoding/json"
	"reflect"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/parser"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
	"yougame.com/yougame-server/util"
)

type ApiWishListController struct {
	api.ApiController
}

func (c *ApiWishListController) GetWishList() {
	var err error
	defer api.CheckError(func(e error) {
		api.HandleApiError(c.Controller, e)
	})
	queryBuilder := service.WishListQueryBuilder{}
	page, pageSize := c.GetPage()
	queryBuilder.WithPage(service.PageOption{
		Page:     page,
		PageSize: pageSize,
	})

	userId, _ := c.GetInt64("user", 0)
	if userId != 0 {
		queryBuilder.BelongToUser(userId)
	}
	count, wishlist, err := queryBuilder.GetWishList()
	if err != nil {
		panic(err)
	}

	results := make([]interface{}, 0)
	for _, item := range wishlist {
		results = append(results, reflect.ValueOf(*item).Interface())
	}
	serializerDataList := serializer.SerializeMultipleData(&serializer.WishListModel{}, results, util.GetSiteAndPortUrl(c.Controller))
	c.ServerPageResult(serializerDataList, count, page, pageSize)
}

func (c *ApiWishListController) DeleteWishListItems() {
	var err error
	defer api.CheckError(func(e error) {
		api.HandleApiError(c.Controller, e)
	})
	claims, err := security.ParseAuthHeader(c.Controller)
	if err != nil {
		panic(security.ReadAuthorizationFailed)
	}
	if claims == nil {
		panic(security.ReadAuthorizationFailed)
	}

	requestStruct := parser.DeleteWishlistItems{}
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestStruct)
	if err != nil {
		panic(api.ParseJsonDataError)
	}
	permission := DeleteWishlistPermission{}
	isAllow := permission.DeleteWishlistPermission(map[string]interface{}{
		"ids":    requestStruct.Items,
		"claims": *claims,
	})
	if !isAllow {
		panic(api.PermissionDeniedError)
	}

	queryBuilder := service.WishListQueryBuilder{}
	queryBuilder.BelongToUser(claims.UserId)
	for _, id := range requestStruct.Items {
		queryBuilder.InId(id)
	}
	err = queryBuilder.DeleteWishList()
	if err != nil {
		panic(err)
	}
	c.Data["json"] = serializer.CommonApiResponseBody{
		Success: true,
	}
	c.ServeJSON()

}
