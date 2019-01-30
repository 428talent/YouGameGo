package cart

import (
	"strconv"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/parser"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
	"yougame.com/yougame-server/util"
)

type ApiCartController struct {
	api.ApiController
}

func (c ApiCartController) GetCartList() {
	c.WithErrorContext(func() {
		listView := api.ListView{
			Controller:    &c.ApiController,
			QueryBuilder:  &service.CartQueryBuilder{},
			ModelTemplate: serializer.NewCartTemplate(serializer.DefaultCartTemplateType),
			Init: func() {
				c.GetAuth()
			},
			SetFilter: func(builder service.ApiQueryBuilder) {
				cartQueryBuilder := builder.(*service.CartQueryBuilder)
				cartQueryBuilder.InUser(c.User.Id)
				cartQueryBuilder.WithEnable("visit")
				util.FilterByParam(&c.Controller, "good", builder, "WithGood", false)
			},
		}
		err := listView.Exec()
		if err != nil {
			panic(err)
		}

	})

}

func (c *ApiCartController) Create() {
	c.WithErrorContext(func() {
		createView := api.CreateView{
			Controller:    &c.ApiController,
			Parser:        &parser.CreateCartRequestBody{},
			Model:         &models.CartItem{},
			ModelTemplate: serializer.NewCartTemplate(serializer.DefaultCartTemplateType),
			OnPrepareSave: func(c *api.CreateView) {
				dataModel := c.Model.(*models.CartItem)
				requestParser := c.Parser.(*parser.CreateCartRequestBody)
				dataModel.Good = &models.Good{
					Id: int(requestParser.GoodId),
				}
				dataModel.User = c.Controller.User
				dataModel.Enable = true
			},
			Validate: func(v *api.CreateView) {
				requestParser := v.Parser.(*parser.CreateCartRequestBody)
				checkDuplicateCartItemValidator := DuplicateCartItemValidator{}
				result := checkDuplicateCartItemValidator.Check(map[string]interface{}{
					"userId": int64(c.User.Id),
					"goodId": requestParser.GoodId,
				})
				if !result {
					panic(api.InvalidateError)
				}
			},
		}
		err := createView.Exec()
		if err != nil {
			panic(err)
		}

	})
}

func (c *ApiCartController) DeleteItem() {
	c.WithErrorContext(func() {
		deleteView := api.DeleteView{
			Controller: &c.ApiController,
			Model:      &models.CartItem{},
			Permissions: []api.PermissionInterface{
				&DeleteCartPermission{},
			},
			GetPermissionContext: func(permissionContext *map[string]interface{}) *map[string]interface{} {
				context := *permissionContext
				idParam := c.Ctx.Input.Param(":id")
				id, err := strconv.Atoi(idParam)
				if err != nil {
					panic(err)
				}
				context["cartId"] = id
				return permissionContext
			},
		}
		err := deleteView.Exec()
		if err != nil {
			panic(err)
		}
	})
}
