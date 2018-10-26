package cart

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"yougame.com/letauth/util"
	"yougame.com/letauthsdk/auth"
	"yougame.com/yougame-server/controllers/web"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
)

type CartController struct {
	beego.Controller
}

func (c *CartController) Get() {
	claims, err := auth.ParseAuthCookie(c.Controller, security.AppSecret)
	if err != nil {
		beego.Error(err)
	}
	web.SetPageAuthInfo(c.Controller, claims)

	page, pageSize := util.ReadPageParam(c.Controller)
	if claims == nil {
		c.Redirect("/login", 302)
		return
	}

	var cartList []*models.CartItem
	cartList, err = models.GetCartList(func(o orm.QuerySeter) orm.QuerySeter {
		return o.Filter("user_id", claims.UserId).Limit(pageSize).Offset((page - 1) * pageSize)
	})
	totalPrice := 0.0
	for _, cartItem := range cartList {
		err = cartItem.ReadGood()
		if err != nil {
			beego.Error(err)
		}
		err = cartItem.Good.ReadGame()
		if err != nil {
			beego.Error(err)
		}
		totalPrice += cartItem.Good.Price
		err = cartItem.Good.ReadGame()
		if err != nil {
			beego.Error(err)
		}
		err = cartItem.Good.Game.ReadGameBand()
		if err != nil {
			beego.Error(err)
		}
	}
	c.Data["CartList"] = cartList
	c.Data["TotalPrice"] = totalPrice
	c.TplName = "cart/index.html"
}

func (c *CartController) Post() {
	claims, err := auth.ParseAuthCookie(c.Controller, security.AppSecret)
	if err != nil {
		beego.Error(err)
	}
	if claims == nil {
		c.Redirect("/login", 302)
		return
	}
	goodId, err := c.GetInt("GoodId")
	if err != nil {
		beego.Error(err)
	}

	cartItem := models.CartItem{
		Good: &models.Good{
			Id: goodId,
		},
		UserId: claims.UserId,
	}
	err = cartItem.Save()
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/cart", 302)
}

func (c *CartController) RemoveCartItem() {
	claims, err := auth.ParseAuthCookie(c.Controller, security.AppSecret)
	if err != nil {
		beego.Error(err)
	}
	if claims == nil {
		c.Redirect("/login", 302)
		return
	}
	cartItemId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		beego.Error(err)
	}
	cartItem := models.CartItem{Id: cartItemId}
	err = cartItem.QueryCartById()
	if err != nil {
		beego.Error(err)
	}
	if cartItem.UserId != claims.UserId {
		c.Redirect("/login", 302)
		return
	}
	err = cartItem.Delete()
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/cart", 302)

}

func (c *CartController)ClearAll() {
	claims, err := auth.ParseAuthCookie(c.Controller, security.AppSecret)
	if err != nil {
		beego.Error(err)
	}
	if claims == nil {
		c.Redirect("/login", 302)
		return
	}
	cartItem := models.CartItem{UserId:claims.UserId}
	err = cartItem.DeleteAll()
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/cart", 302)
}