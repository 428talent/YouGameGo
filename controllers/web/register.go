package web

import (
	"github.com/astaxie/beego"
	"yougame.com/yougame-server/security"
)

type RegisterController struct {
	WebController
}

func (c *RegisterController) Get() {
	claims, err := security.ParseAuthCookies(c.Controller)
	if err != nil {
		beego.Error(err)
	}

	c.SetPageAuthInfo(claims)
	c.TplName = "register.tpl"
}
