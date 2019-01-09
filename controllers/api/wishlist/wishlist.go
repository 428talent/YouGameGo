package wishlist

import (
	"encoding/json"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
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
	queryBuilder.OnlyEnable(true)
	count, wishlist, err := queryBuilder.GetWishList()
	if err != nil {
		panic(err)
	}
	serializerDataList := serializer.SerializeMultipleTemplate(wishlist, serializer.NewWishlistTemplate(serializer.DefaultCartTemplateType), map[string]interface{}{
		"site": util.GetSiteAndPortUrl(c.Controller),
	})
	c.ServerPageResult(serializerDataList, count, page, pageSize)
}
func (c *ApiWishListController) Create() {
	createView := api.CreateView{
		Controller: &c.ApiController,
		Parser:     &parser.CreateWishlistRequestBody{},
		Model:      &models.WishList{},
		ModelTemplate:serializer.NewWishlistTemplate(serializer.DefaultWishListTemplateType),
		OnPrepareSave: func(c *api.CreateView) {
			model := c.Model.(*models.WishList)
			parserModel := c.Parser.(*parser.CreateWishlistRequestBody)
			model.Game = &models.Game{
				Id: int(parserModel.GameId),
			}
			model.UserId = c.Controller.User.Id
		},
	}
	err := createView.Exec()
	if err != nil {
		panic(err)
	}

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
