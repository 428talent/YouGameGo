package order

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"yougame.com/letauthsdk/auth"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/parser"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/serializer"
)

type ApiOrderController struct {
	beego.Controller
}

func (c *ApiOrderController) CreateOrder() {
	claims, err := auth.ParseAuthHeader(c.Controller, security.AppSecret)
	if err != nil {
		beego.Error(err)
		return
	}
	if claims == nil {
		return
	}
	requestBodyStruct := parser.CreateOrderParser{}
	err = requestBodyStruct.Parse(c.Controller)
	if err != nil {
		beego.Error(err)
		return
	}
	var goodList []*models.OrderGood
	for _, goodId := range requestBodyStruct.Goods {
		goodList = append(goodList, &models.OrderGood{
			Good: &models.Good{
				Id: int(goodId),
			},
		})
	}
	order := models.Order{
		UserId: claims.UserId,
		State:  "Created",
		Goods:  goodList,
	}
	err = order.SaveOrder()
	if err != nil {
		beego.Error(err)
	}
	c.ServeJSON()
}

func (c *ApiOrderController) GetOrderList() {
	claims, err := auth.ParseAuthHeader(c.Controller, security.AppSecret)
	if err != nil {
		beego.Error(err)
		return
	}
	if claims == nil {
		return
	}
	orders,err := models.GetOrderList(func(o orm.QuerySeter) orm.QuerySeter {
		return o.Filter("user_id",claims.Id)
	})
	if err != nil {
		beego.Error(err)
	}
	serializedData,err := serializer.SerializeOrderList(orders,serializer.OrderSerializer{})
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = serializedData
	c.ServeJSON()
}
