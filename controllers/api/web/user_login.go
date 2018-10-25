package api_web

import (
	"github.com/astaxie/beego"
	"yougame.com/letauthsdk/client"
	"yougame.com/yougame-server/auth"
)

type UserLoginController struct {
	beego.Controller
}

func (c *UserLoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	authReqeustBody := client.LoginRequestBody{
		LoginName: username,
		Password: password,
	}
	response,err := auth.AuthClient.Login(authReqeustBody)
	if err != nil {
		beego.Error(err)
	}
	if response.Success {
		beego.Debug(response.Sign)
		c.Ctx.SetCookie("yougame_token", response.Sign)
		c.Redirect("/", 302)
	}else{
		c.Redirect("/login", 302)
	}
}
