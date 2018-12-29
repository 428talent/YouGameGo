package good

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
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
