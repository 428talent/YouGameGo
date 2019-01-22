package tag

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

func (c *Controller) CreateTag() {
	c.WithErrorContext(func() {
		view := api.CreateView{
			Controller:    &c.ApiController,
			Parser:        &parser.CreateTagRequestBody{},
			ModelTemplate: serializer.NewTagTemplate(serializer.DefaultTagTemplateType),
			Model:         &models.Tag{},
		}
		err := view.Exec()
		if err != nil {
			panic(err)
		}
	})
}

func (c *Controller) List() {
	c.WithErrorContext(func() {
		view := api.ListView{
			Controller:    &c.ApiController,
			ModelTemplate: serializer.NewTagTemplate(serializer.DefaultTagTemplateType),
			QueryBuilder:  &service.TagQueryBuilder{},
			SetFilter: func(builder service.ApiQueryBuilder) {
				util.FilterByParam(&c.Controller, "id", builder, "InId")
				util.FilterByParam(&c.Controller, "name", builder, "WithName")
			},
		}
		err := view.Exec()
		if err != nil {
			panic(err)
		}
	})
}

func (c *Controller) Update() {
	c.WithErrorContext(func() {
		view := api.UpdateView{
			Controller:    &c.ApiController,
			ModelTemplate: serializer.NewTagTemplate(serializer.DefaultTagTemplateType),
			Parser:        &parser.CreateTagRequestBody{},
			Model:         &models.Tag{},
		}
		err := view.Exec()
		if err != nil {
			panic(err)
		}
	})
}

func (c *Controller) DeleteTag() {
	c.WithErrorContext(func() {
		view := api.DeleteView{
			Controller:    &c.ApiController,
			Model:         &models.Tag{},
		}
		err := view.Exec()
		if err != nil {
			panic(err)
		}
	})
}