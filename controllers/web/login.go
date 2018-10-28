package web

import (
	"github.com/astaxie/beego"
	"yougame.com/letauthsdk/client"
	"yougame.com/yougame-server/auth"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
)

type AuthorizationController struct {
	beego.Controller
}

func (c *AuthorizationController) Get() {
	c.TplName = "login.tpl"
}

func (c *AuthorizationController) Logout() {
	c.Ctx.SetCookie("yougame_token", "", 1)
	c.Redirect("/", 302)
}

func (c *AuthorizationController) Login() {
	username := c.GetString("username")
	password := c.GetString("password")
	user := models.User{Username: username, Password: password}
	if !models.CheckUserValidate(&user) {
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

func (c *AuthorizationController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")

	authBody, err := auth.AuthClient.CreateUser(client.CreateUserRequestBody{
		Username: username,
		Password: password,
	})
	if err != nil {
		beego.Error(err)
	}
	beego.Debug(authBody.Success)
	c.Redirect("/login", 302)
}