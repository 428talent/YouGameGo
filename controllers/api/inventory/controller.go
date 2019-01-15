package inventory

import (
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
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
				inventoryQueryBuilder.BelongUser(c.User.Id)

				for _, goodIdParam := range c.GetStrings("good") {
					inventoryQueryBuilder.InGood(goodIdParam)
				}
			},
		}
		err := listView.Exec()
		if err != nil {
			panic(err)
		}

	})
}
