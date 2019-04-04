package inventory

import (
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
	"yougame.com/yougame-server/util"
)

type Controller struct {
	api.ApiController
}

func (c *Controller) GetInventoryList() {
	c.WithErrorContext(func() {
		listView := api.ListView{
			Init: func() {
				c.GetAuth()
			},
			Controller:    &c.ApiController,
			QueryBuilder:  &service.InventoryQueryBuilder{},
			ModelTemplate: serializer.NewInventoryTemplate(serializer.DefaultInventoryTemplateType),
			SetFilter: func(builder service.ApiQueryBuilder) {
				inventoryQueryBuilder := builder.(*service.InventoryQueryBuilder)
				if len(c.Controller.Ctx.Input.Param("user")) > 0 {
					util.FilterByParam(&c.Controller, "good", builder, "BelongUser", false)
				} else {
					inventoryQueryBuilder.BelongUser(c.User.Id)
				}
				util.FilterByParam(&c.Controller, "good", builder, "InGood", false)
				util.FilterByParam(&c.Controller, "game", builder, "InGame", false)
			},
		}
		err := listView.Exec()
		if err != nil {
			panic(err)
		}

	})
}
