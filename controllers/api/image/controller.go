package image

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

func (c *Controller) DeleteImage() {
	c.WithErrorContext(func() {
		deleteView := api.DeleteView{
			Controller: &c.ApiController,
			Permissions: []api.PermissionInterface{
				&DeleteImagePermission{},
			},
			Model: &models.Image{},
		}
		err := deleteView.Exec()
		if err != nil {
			panic(err)
		}
	})
}

func (c *Controller) UpdateImage() {
	c.WithErrorContext(func() {
		updateView := api.UpdateView{
			Controller: &c.ApiController,
			Parser:     &parser.UpdateImageRequestBody{},
			Permissions: []api.PermissionInterface{
				&UpdateImagePermission{},
			},
			ModelTemplate: &serializer.ImageAdminTemplate{},
			Model:         &models.Image{},
		}
		err := updateView.Exec()
		if err != nil {
			panic(err)
		}
	})
}

func (c *Controller) GetImageList() {
	c.WithErrorContext(func() {
		listView := api.ListView{
			Controller:    &c.ApiController,
			QueryBuilder:  &service.ImageQueryBuilder{},
			ModelTemplate: &serializer.ImageAdminTemplate{},
			SetFilter: func(builder service.ApiQueryBuilder) {
				util.FilterByParam(&c.Controller, "name", builder, "WithName", false)
			},
		}
		err := listView.Exec()
		if err != nil {
			panic(err)
		}
	})
}
