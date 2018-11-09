package web

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/service"
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

	//获取评论
	comments, err := models.GetCommentList(func(o orm.QuerySeter) orm.QuerySeter {
		return o.Filter("Good__Game__Id", gameId).Filter("Enable", true).Limit(30)
	})
	if err != nil {
		panic(err)
	}
	var positiveComment []*models.Comment
	var negativeComment []*models.Comment
	for _, comment := range comments {
		for _, gameGood := range game.Goods {
			if gameGood.Id == comment.Good.Id {
				comment.Good = gameGood
			}
		}
		commentUser, err := models.GetUserById(comment.User.Id)
		if err != nil {
			continue
		}
		comment.User = commentUser
		err = comment.User.ReadProfile()
		if err != nil {
			continue
		}
		if comment.Evaluation == models.EvaluationPositive {
			positiveComment = append(positiveComment, comment)
		} else if comment.Evaluation == models.EvaluationNegative {
			negativeComment = append(negativeComment, comment)
		}

	}
	totalComment, positiveCommentCount, negativeCommentCount, err := service.GetGameCommentStatistics(game.Id)
	if err != nil {
		panic(err)
	}
	if totalComment != 0 {
		c.Data["commentPositiveRate"] = fmt.Sprintf("%.2f", float64(positiveCommentCount)/float64(totalComment) * 100)
	}
	c.Data["commentCount"] = totalComment
	c.Data["commentPositiveCount"] = positiveCommentCount
	c.Data["commentNegativeCount"] = negativeCommentCount

	c.Data["positiveComment"] = positiveComment
	c.Data["negativeComment"] = negativeComment
	c.Data["GameGoods"] = gameGoods
	c.Data["game"] = game
	c.TplName = "detail/detail.html"
}
