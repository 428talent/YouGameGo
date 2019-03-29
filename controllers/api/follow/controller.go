package follow

import (
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/parser"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
	"yougame.com/yougame-server/util"
)

type Controller struct {
	api.ApiController
}

func (c *Controller) Create() {
	c.WithErrorContext(func() {
		createView := api.CreateView{
			Controller:    &c.ApiController,
			Parser:        &parser.CreateFollowRequestBody{},
			Model:         &models.Follow{},
			ModelTemplate: serializer.NewFollowSerializeTemplate(serializer.DefaultFollowTemplateType),
			OnPrepareSave: func(c *api.CreateView) {
				model := c.Model.(*models.Follow)
				requestParser := c.Parser.(*parser.CreateFollowRequestBody)
				model.User = &models.User{
					Id: requestParser.UserId,
				}
				model.Following = &models.User{
					Id: requestParser.FollowingId,
				}
				model.Enable = true
			},
			Validators: []api.RequestValidator{
				&DuplicateFollowValidator{},
			},
		}
		err := createView.Exec()
		if err != nil {
			panic(err)
		}
	})
}

func (c *Controller) List() {
	queryBuilder := &service.FollowQueryBuilder{}
	c.WithErrorContext(func() {
		view := api.ListView{
			Controller:    &c.ApiController,
			QueryBuilder:  queryBuilder,
			ModelTemplate: serializer.NewFollowSerializeTemplate(serializer.DefaultFollowTemplateType),
			SetFilter: func(builder service.ApiQueryBuilder) {
				util.FilterByParam(&c.Controller, "user", builder, "InUser", false)
				util.FilterByParam(&c.Controller, "following", builder, "InFollowing", false)
				util.FilterByParam(&c.Controller, "id", builder, "InId", false)
				util.FilterByParam(&c.Controller, "order", builder, "ByOrder", false)
			},
		}
		err := view.Exec()
		if err != nil {
			panic(err)
		}
	})
}

func (c *Controller) Delete() {
	c.WithErrorContext(func() {
		view := api.DeleteView{
			Controller: &c.ApiController,
			Model:      &models.Follow{},
		}
		err := view.Exec()
		if err != nil {
			panic(err)
		}
	})
}

