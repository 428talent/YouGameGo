package web

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"yougame.com/yougame-server/forms"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/service"
)

type CommentController struct {
	WebController
}

func (c *CommentController) WriteComment() {
	claims, err := security.ParseAuthCookies(c.Controller)
	if err != nil {
		beego.Error(err)
	}
	c.SetPageAuthInfo(claims)
	goodId, err := strconv.Atoi(c.GetString(":goodId"))
	if err != nil {
		beego.Error(err)
	}
	good, err := service.GetGoodById(goodId)
	if err != nil {
		beego.Error(err)
	}
	err = good.ReadGame()
	if err != nil {
		beego.Error(err)
	}
	err = good.Game.ReadGameBand()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Good"] = good
	c.TplName = "comment/write.html"
}

func (c *CommentController) SaveComment() {
	claims, err := security.ParseAuthCookies(c.Controller)
	if err != nil {
		beego.Error(err)
	}
	c.LoadRequestUser(claims)
	form := forms.CreateCommentForm{}
	err = c.ParseForm(&form)
	if err != nil {
		beego.Error(err)
		return
	}

	if err != nil {
		beego.Error(errors.New("permission denied"))
		return
	}

	good, _, err := service.CreateGoodComment(c.User, form.Content, form.GoodId, form.Evaluation)
	if err != nil {
		beego.Error(err)
		return
	}

	c.Redirect(fmt.Sprintf("/game/%d/", good.Game.Id), 302)

}
