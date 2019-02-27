package permission

import (
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
	"yougame.com/yougame-server/util"
)

type Controller struct {
	api.ApiController
}

func (c *Controller) List() {
	c.WithErrorContext(func() {
		listView := api.ListView{
			Controller:    &c.ApiController,
			QueryBuilder:  &service.PermissionQueryBuilder{},
			ModelTemplate: serializer.NewPermissionTemplate(serializer.DefaultPermissionTemplateType),
			SetFilter: func(builder service.ApiQueryBuilder) {
				util.FilterByParam(&c.Controller, "userGroup", builder, "WithUserGroup", false)
				util.FilterByParam(&c.Controller, "name", builder, "WithName", true)
			},
		}
		err := listView.Exec()
		if err != nil {
			panic(err)
		}
	})
}
