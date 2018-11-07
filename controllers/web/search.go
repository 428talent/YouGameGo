package web

import (
	"github.com/astaxie/beego"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
)

type SearchController struct {
	WebController
}

func (c *SearchController) Get() {
	claims, err := security.ParseAuthCookies(c.Controller)
	if err != nil {
		beego.Error(err)
	}

	c.SetPageAuthInfo(claims)
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
