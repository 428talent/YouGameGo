package web

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"yougame.com/letauthsdk/auth"
	"yougame.com/yougame-server/forms"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/service"
)

type OrderController struct {
	beego.Controller
}

func (c *OrderController) Get() {
	claims, err := auth.ParseAuthCookie(c.Controller, security.AppSecret)
	if err != nil {
		beego.Error(err)
	}

	SetPageAuthInfo(c.Controller, claims)

	if claims == nil {
		c.Redirect("/login", 302)
		return
	}

	orderId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		beego.Error(err)
		return
	}
	order := models.Order{Id: orderId}
	err = order.QueryById()
	if err != nil {
		beego.Error(err)
	}
	err = order.ReadOrderGoods()
	if err != nil {
		beego.Error(err)
	}
	totalPrice := 0.0
	for _, orderGood := range order.Goods {
		err = orderGood.QueryById()
		if err != nil {
			beego.Error(err)
		}
		totalPrice += orderGood.Price
		err = orderGood.ReadGood()
		if err != nil {
			beego.Error(err)
		}
		err = orderGood.Good.ReadGame()
		if err != nil {
			beego.Error(err)
		}
		err = orderGood.Good.Game.ReadGameBand()
		if err != nil {
			beego.Error(err)
		}
	}
	c.Data["IsPaid"] = order.State == "Done"
	c.Data["Order"] = order
	c.Data["TotalPrice"] = totalPrice
	c.TplName = "order/index.html"
}

func (c *OrderController) CreateOrder() {
	claims, err := auth.ParseAuthCookie(c.Controller, security.AppSecret)
	if err != nil {
		beego.Error(err)
	}

	if claims == nil {
		c.Redirect("/login", 302)
		return
	}

	cartList, err := models.GetCartList(func(o orm.QuerySeter) orm.QuerySeter {
		return o.Filter("user_id", claims.UserId)
	})
	if err != nil {
		beego.Error(err)
	}
	order := models.Order{
		State:  "Created",
		User: &models.User{Id:claims.UserId,},
	}
	var orderGoodList []*models.OrderGood
	for _, cartItem := range cartList {
		orderGoodList = append(orderGoodList, &models.OrderGood{Good: cartItem.Good})
	}
	order.Goods = orderGoodList
	err = order.SaveOrder()
	if err != nil {
		beego.Error(err)
	}
	c.Redirect(fmt.Sprintf("/order/%d", order.Id), 302)
}

func (c *OrderController) PayOrder() {
	claims, err := auth.ParseAuthCookie(c.Controller, security.AppSecret)
	if err != nil {
		beego.Error(err)
	}

	if claims == nil {
		c.Redirect("/login", 302)
		return
	}

	form := forms.PayOrderForm{}
	err = c.ParseForm(&form)
	if err != nil {
		beego.Error(err)
	}
	order := models.Order{Id:form.OrderId}
	err = order.QueryById()
	if err != nil {
		beego.Error(err)
	}
	err = service.PayOrder(order)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect(fmt.Sprintf("/order/%d", order.Id), 302)
}
