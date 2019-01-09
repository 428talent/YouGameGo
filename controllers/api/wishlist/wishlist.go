package wishlist

import (
	"encoding/json"
	"strconv"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/parser"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
)

type ApiWishListController struct {
	api.ApiController
}

func (c *ApiWishListController) GetWishList() {
	c.WithErrorContext(func() {
		listView := api.ListView{
			Controller: &c.ApiController,
			Init: func() {
				c.GetAuth()
			},
			QueryBuilder:  &service.WishListQueryBuilder{},
			ModelTemplate: serializer.NewWishlistTemplate(serializer.DefaultWishListTemplateType),
			SetFilter: func(builder service.ApiQueryBuilder) {
				wishlistItemQueryBuilder := builder.(*service.WishListQueryBuilder)
				wishlistItemQueryBuilder.BelongToUser(c.User.Id)
				wishlistItemQueryBuilder.WithEnable("visit")
				gameParamList := c.GetStrings("game")
				for _, gameParam := range gameParamList {
					gameId, err := strconv.Atoi(gameParam)
					if err != nil {
						panic(err)
					}
					wishlistItemQueryBuilder.WithGame(gameId)
				}
			},
		}
		err := listView.Exec()
		if err != nil {
			panic(err)
		}
	})

}
func (c *ApiWishListController) Create() {
	c.WithErrorContext(func() {
		createView := api.CreateView{
			Controller:    &c.ApiController,
			Parser:        &parser.CreateWishlistRequestBody{},
			Model:         &models.WishList{},
			ModelTemplate: serializer.NewWishlistTemplate(serializer.DefaultWishListTemplateType),
			OnPrepareSave: func(c *api.CreateView) {
				model := c.Model.(*models.WishList)
				parserModel := c.Parser.(*parser.CreateWishlistRequestBody)
				model.Game = &models.Game{
					Id: int(parserModel.GameId),
				}
				model.UserId = c.Controller.User.Id
				model.Enable = true

				queryBuilder := service.WishListQueryBuilder{}
				queryBuilder.BelongToUser(c.Controller.User.Id)
				queryBuilder.WithGame(parserModel.GameId)
				queryBuilder.WithEnable("visit")
				count, _, err := queryBuilder.GetWishList()
				if err != nil {
					panic(err)
				}
				if count > 0 {
					panic(api.DuplicateResourceError)
				}
			},
		}
		err := createView.Exec()
		if err != nil {
			panic(err)
		}
	})

}

func (c *ApiWishListController) DeleteItem() {
	c.WithErrorContext(func() {
		deleteView := api.DeleteView{
			Controller:  &c.ApiController,
			Model:       &models.WishList{},
			Permissions: []api.PermissionInterface{},
			GetPermissionContext: func(permissionContext *map[string]interface{}) *map[string]interface{} {
				idParam := c.Ctx.Input.Param(":id")
				id, err := strconv.Atoi(idParam)
				if err != nil {
					panic(err)
				}
				(*permissionContext)["id"] = id
				return permissionContext
			},
		}
		err := deleteView.Exec()
		if err != nil {
			panic(err)
		}

	})
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
