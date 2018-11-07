package web

import (
	"github.com/astaxie/beego"
	"yougame.com/yougame-server/security"
)

type CommentController struct {
	WebController
}

func (c *CommentController)WriteComment(){
	claims, err := security.ParseAuthCookies(c.Controller)
	if err != nil {
		beego.Error(err)
	}
	c.SetPageAuthInfo(claims)
	c.TplName = "comment/write.html"
}

