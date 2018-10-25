package web

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"yougame.com/letauthsdk/auth"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	claims, err := auth.ParseAuthCookie(c.Controller, security.AppSecret)
	if err != nil {
		beego.Error(err)
	}

	SetPageAuthInfo(c.Controller, claims)

	gameList, err := models.GetGameList(func(o orm.QuerySeter) orm.QuerySeter {
		return o.Limit(8)
	})
	if err != nil {
		beego.Error(err)
	}
	for _, e := range gameList {
		err = e.ReadGameBand()
	}
	c.Data["GameList"] = gameList
	c.TplName = "index.tpl"
}
