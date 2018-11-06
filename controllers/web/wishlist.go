package web

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"yougame.com/letauthsdk/auth"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
)

type WishListController struct {
	beego.Controller
}

func (c *WishListController) SaveWishList() {
	claims, err := auth.ParseAuthCookie(c.Controller, security.AppSecret)
	if err != nil {
		beego.Error(err)
	}

	if claims == nil {
		beego.Error(errors.New("请先登录后操作"))
	}

	gameId, err := c.GetInt("GameId")
	if err != nil {
		beego.Error(err)
	}
	wishList := models.WishList{
		UserId: claims.UserId,
		Game: &models.Game{
			Id: gameId,
		},
	}
	err = models.SaveWishList(&wishList)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect(fmt.Sprintf("/game/%d", gameId), 302)
}
