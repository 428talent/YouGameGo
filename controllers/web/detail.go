package web

import (
	"github.com/astaxie/beego"
	"strconv"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
)

type DetailController struct {
	WebController
}

func (c *DetailController) Get() {
	claims, err := security.ParseAuthCookies(c.Controller)
	if err != nil {
		beego.Error(err)
	}

	c.SetPageAuthInfo(claims)

	gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		beego.Error(err)
		return
	}

	game := models.Game{Id: gameId}
	err = game.QueryById()
	if err != nil {
		beego.Error(err)
		return
	}

	game.ReadGameBand()
	err = game.ReadGamePreviewImage()
	if err != nil {
		beego.Error(err)
	}

	err = game.ReadTags()
	if err != nil {
		beego.Error(err)
	}

	err = game.ReadGoods()
	if err != nil {
		beego.Error(err)
	}

	beego.Debug(game)
	c.Data["game"] = game
	c.TplName = "detail/detail.html"
}
