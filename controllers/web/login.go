package web

import (
	"github.com/astaxie/beego"
	"yougame.com/letauthsdk/auth"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
)

type AuthorizationController struct {
	beego.Controller
}

func (c *AuthorizationController) Get() {
	claims, err := auth.ParseAuthCookie(c.Controller, security.AppSecret)
	if err != nil {
		beego.Error(err)
	}

	SetPageAuthInfo(c.Controller,claims)
	c.TplName = "login.tpl"
}

func (c *AuthorizationController) Logout() {
	c.Ctx.SetCookie("yougame_token", "", 1)
	c.Redirect("/", 302)
}

func (c *AuthorizationController) Login() {
	username := c.GetString("username")
	password := c.GetString("password")
	flash := beego.NewFlash()
	user := models.User{Username: username, Password: password}
	if !models.CheckUserValidate(&user) {
		flash.Set("ErrorTitle","登录错误")
		flash.Set("ErrorContent", "请检查用户名和密码是否正确")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}
	signString, err := security.GenerateJWTSign(&user)
	if err != nil {
		beego.Error(err)
		return
	}
	c.Ctx.SetCookie("yougame_token", *signString)
	c.Redirect("/", 302)

}

func (c *AuthorizationController) CreateUser() {
	username := c.GetString("username")
	password := c.GetString("password")

	_, err := models.CreateUserAccount(username, password)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/login", 302)
}
