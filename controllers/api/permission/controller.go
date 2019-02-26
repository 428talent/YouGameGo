package permission

import (
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
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
		}
		err := listView.Exec()
		if err != nil {
			panic(err)
		}
	})
}