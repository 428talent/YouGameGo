package web

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.tpl"
}

func (c *LoginController) Logout() {
	c.Ctx.SetCookie("yougame_token", "", 1)
	c.Redirect("/", 302)
}
