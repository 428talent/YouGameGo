package profile

import (
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
)

type Controller struct {
	api.ApiController
}

func (c *Controller) GetProfileList() {
	c.WithErrorContext(func() {
		listView := api.ListView{
			Controller:    &c.ApiController,
			QueryBuilder:  &service.UserProfileQueryBuilder{},
			ModelTemplate: serializer.NewProfileTemplate(serializer.DefaultProfileTemplateType),
			SetFilter: func(builder service.ApiQueryBuilder) {
				profileQueryBuilder := builder.(*service.UserProfileQueryBuilder)
				for _, userId := range c.GetStrings("user") {
					profileQueryBuilder.InUser(userId)
				}
			},
			OnGetResult: func(i interface{}) {
				profiles := i.([]*models.Profile)
				for _, profile := range profiles {
					profile.ReadUser()
				}
			},
		}
		err := listView.Exec()
		if err != nil {
			panic(err)
		}
	})
}
