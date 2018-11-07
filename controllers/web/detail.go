package web

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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

	userBoughtGood, err := models.GetGoodList(func(o orm.QuerySeter) orm.QuerySeter {
		o.Filter("Users__User__Id", c.User.Id).Filter("game__id", gameId)
		return o
	})
	type GameGood struct {
		Good   models.Good
		Bought bool
	}
	var gameGoods []GameGood
	for _, gameGood := range game.Goods {
		item := GameGood{
			Good:   *gameGood,
			Bought: false,
		}

		for _, boughtGood := range userBoughtGood {
			if boughtGood.Id == gameGood.Id {
				item.Bought = true
			}
		}
		gameGoods = append(gameGoods, item)
	}

	beego.Debug(game)
	c.Data["GameGoods"] = gameGoods
	c.Data["game"] = game
	c.TplName = "detail/detail.html"
}
