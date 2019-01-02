package good

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/parser"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
	"yougame.com/yougame-server/util"
)

type Controller struct {
	api.ApiController
}

func (c *Controller) UpdateGood() {
	c.WithErrorContext(func() {
		goodId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
		if err != nil {
			panic(err)
		}

		claims, err := c.GetAuth()
		if err != nil {
			panic(api.ClaimsNoFoundError)
		}
		if claims == nil {
			panic(api.ClaimsNoFoundError)
		}

		err = c.CheckPermission([]api.PermissionInterface{
			&UpdateGoodPermission{},
		}, map[string]interface{}{
			"claims": claims,
		})
		if err != nil {
			panic(err)
		}

		requestBody := &parser.UpdateGoodRequestBody{}
		err = json.Unmarshal(c.Ctx.Input.RequestBody, requestBody)

		good := &models.Good{
			Id:    goodId,
			Name:  requestBody.Name,
			Price: requestBody.Price,
		}

		updateFields := util.GetUpdateModelField(requestBody)
		beego.Debug(updateFields)
		err = service.UpdateGood(good, updateFields...)
		if err != nil {
			panic(err)
		}

		goodTemplate := serializer.NewGoodSerializeTemplate(serializer.AdminGoodTemplateType)
		goodTemplate.Serialize(good, map[string]interface{}{
			"site": util.GetSiteAndPortUrl(c.Controller),
		})
		c.Data["json"] = goodTemplate
		c.ServeJSON()

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
			Controller:    &c.ApiController,
			QueryBuilder:  &service.GoodQueryBuilder{},
			ModelTemplate: serializer.NewGoodSerializeTemplate(serializer.AdminGoodTemplateType),
			GetTemplate: func() serializer.Template {
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

			},
		}
		err := listView.Exec()
		if err != nil {
			panic(err)
		}
	})
}
