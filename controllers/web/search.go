package web

import (
	"github.com/astaxie/beego"
	"yougame.com/letauthsdk/auth"
	AppAuth "yougame.com/yougame-server/auth"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
)

type SearchController struct {
	beego.Controller
}

func (c *SearchController) Get() {
	claims, err := auth.ParseAuthCookie(c.Controller, security.AppSecret)
	if err != nil {
		beego.Error(err)
	}

	if claims != nil {
		user, err := AppAuth.AuthClient.GetUser(claims.UserId)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Nickname"] = user.Profile.Nickname
		if len(user.Profile.Avatar) == 0 {
			c.Data["Avatar"] = "/static/img/user.png"
		} else {
			c.Data["Avatar"] = user.Profile.Avatar
		}
	}

	key := c.Ctx.Input.Param(":key")
	gameList, err := models.SearchGame(key)
	if err != nil {
		beego.Error(err)
	}
	for _, game := range gameList {
		err = game.ReadGameBand()
		if err != nil {
			beego.Error(err)
		}
	}
	c.Data["results"] = gameList
	c.TplName = "search/search.html"
}