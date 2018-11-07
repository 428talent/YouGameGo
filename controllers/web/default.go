package web

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
)

type MainController struct {
	WebController
}

func (c *MainController) Get() {
	claims, err := security.ParseAuthCookies(c.Controller)
	if err != nil {
		beego.Error(err)
	}

	c.SetPageAuthInfo( claims)

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
