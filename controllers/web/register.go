package web

import (
	"github.com/astaxie/beego"
	"yougame.com/letauthsdk/auth"
	"yougame.com/yougame-server/security"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	claims, err := auth.ParseAuthCookie(c.Controller, security.AppSecret)
	if err != nil {
		beego.Error(err)
	}

	SetPageAuthInfo(c.Controller,claims)
	c.TplName = "register.tpl"
}
