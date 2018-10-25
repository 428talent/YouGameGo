package api_web

import (
	"github.com/astaxie/beego"
	"yougame.com/letauthsdk/client"
	"yougame.com/yougame-server/auth"
)

type CreateUserController struct {
	beego.Controller
}

func (c *CreateUserController) Post() {
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
