package good

import (
	"github.com/astaxie/beego"
	"strconv"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/parser"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
)

type Controller struct {
	api.ApiController
}

func (c *Controller) UpdateGood() {
	c.WithErrorContext(func() {
		updateView := api.UpdateView{
			Controller: &c.ApiController,
			Parser:     &parser.UpdateGoodRequestBody{},
			Model:      &models.Good{},
			Permissions: []api.PermissionInterface{
				&UpdateGoodPermission{},
			},
			ModelTemplate: serializer.NewGoodSerializeTemplate(serializer.AdminGoodTemplateType),
		}
		err := updateView.Exec()
		if err != nil {
			panic(err)
		}

	})
}
func (c *Controller) GetGood() {
	c.WithErrorContext(func() {
		objectView := api.ObjectView{
			Controller:   &c.ApiController,
			QueryBuilder: &service.GoodQueryBuilder{},
			GetTemplate: func() serializer.Template {
				c.Role = security.Anonymous
				c.GetAuth()
				if c.User != nil {
					if security.CheckUserGroup(c.User, security.UserGroupAdmin) {
						c.Role = security.UserGroupAdmin
					}
				}

				switch c.Role {
				case security.UserGroupAdmin:
					return serializer.NewGoodSerializeTemplate(serializer.AdminGoodTemplateType)
				default:
					return serializer.NewGoodSerializeTemplate(serializer.DefaultGoodTemplateType)
				}

			},
			OnGetResult: func(model interface{}) {
				data := model.(*models.Good)
				beego.Debug(data)
			},
		}
		err := objectView.Exec()
		if err != nil {
			panic(err)
		}
	})
}
func (c *Controller) GetGoods() {
	c.WithErrorContext(func() {
		listView := api.ListView{
			Controller:   &c.ApiController,
			QueryBuilder: &service.GoodQueryBuilder{},
			Init: func() {
				c.GetAuth()
				c.Role = security.Anonymous
				if security.CheckUserGroup(c.User, security.UserGroupAdmin) {
					c.Role = security.UserGroupAdmin
				}
			},
			ModelTemplate: serializer.NewGoodSerializeTemplate(serializer.DefaultGoodTemplateType),
			GetTemplate: func() serializer.Template {
				if c.User != nil && security.CheckUserGroup(c.User, security.UserGroupAdmin) {
					return serializer.NewGoodSerializeTemplate(serializer.AdminGoodTemplateType)
				}
				return serializer.NewGoodSerializeTemplate(serializer.DefaultGoodTemplateType)
			},
			SetFilter: func(builder service.ApiQueryBuilder) {
				goodQueryBuilder, _ := builder.(*service.GoodQueryBuilder)

				idFilters := c.GetStrings("id")
				for _, idParam := range idFilters {
					id, err := strconv.Atoi(idParam)
					if err != nil {
						panic(err)
					}
					goodQueryBuilder.InId(id)
				}

				gameFilters := c.GetStrings("game")
				for _, gameParam := range gameFilters {
					gameId, err := strconv.Atoi(gameParam)
					if err != nil {
						panic(err)
					}
					goodQueryBuilder.InGameId(gameId)
				}

				orders := c.GetStrings("order")
				if len(orders) > 0 {
					goodQueryBuilder.ByOrder(orders...)
				}

				if security.CheckUserGroup(c.User, security.UserGroupAdmin) {
					enable := c.GetString("enable", "visit")
					if enable != "all" {
						goodQueryBuilder.WithEnable(enable)
					}
				}

			},
		}
		err := listView.Exec()
		if err != nil {
			panic(err)
		}
	})
}
func (c *Controller) CreateGood() {
	c.WithErrorContext(func() {
		createView := api.CreateView{
			Controller: &c.ApiController,
			Parser:     &parser.CreateGoodRequestBody{},
			Permissions: []api.PermissionInterface{
				&CreateGoodPermission{},
			},
			Model:         &models.Good{},
			ModelTemplate: serializer.NewGoodSerializeTemplate(serializer.AdminGoodTemplateType),
			OnPrepareSave: func(c *api.CreateView) {
				model := c.Model.(*models.Good)
				requestParser := c.Parser.(*parser.CreateGoodRequestBody)
				model.Game = &models.Game{
					Id: int(requestParser.GameId),
				}
			},
		}
		err := createView.Exec()
		if err != nil {
			panic(err)
		}
	})
}
func (c *Controller) DeleteGood() {
	c.WithErrorContext(func() {
		deleteView := api.DeleteView{
			Controller: &c.ApiController,
			Model:      &models.Good{},
			Permissions: []api.PermissionInterface{
				&DeleteGoodPermission{},
			},
		}
		err := deleteView.Exec()
		if err != nil {
			panic(err)
		}
	})
}
