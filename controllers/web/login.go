package web

import (
	"github.com/astaxie/beego"
	"yougame.com/yougame-server/forms"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/service"
)

type AuthorizationController struct {
	WebController
}

func (c *AuthorizationController) Get() {
	claims, err := security.ParseAuthCookies(c.Controller)
	if err != nil {
		beego.Error(err)
	}

	c.SetPageAuthInfo(claims)
	beego.ReadFromRequest(&c.Controller)
	c.TplName = "login.tpl"
}

func (c *AuthorizationController) Logout() {
	c.Ctx.SetCookie("yougame_token", "", 1)
	c.Redirect("/", 302)
}

func (c *AuthorizationController) Login() {
	form := forms.UserLoginForm{}
	err := c.ParseForm(&form)
	if err != nil {
		beego.Error(err)
	}
	flash := beego.NewFlash()
	jwtString, _, err := service.UserLogin(form.Username, form.Password)
	if err != nil {
		if err == service.LoginUserFailed {
			flash.Set("ErrorTitle", "登录错误")
			flash.Set("ErrorContent", "请检查用户名和密码是否正确")
			flash.Store(&c.Controller)
			c.Redirect("/login", 302)
			return
		}
		beego.Error(err)
		return
	}

	c.Ctx.SetCookie("yougame_token", jwtString)
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
