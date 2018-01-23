package api_web

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"you_game_go/models"
	"you_game_go/util"
)

type CreateUserController struct {
	beego.Controller
}

func (c *CreateUserController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	appConfig, err := config.NewConfig("ini", "./conf/app_local.conf")
	if err != nil {
		beego.Error(err)
	}
	enPassword := util.EncryptSha1(password + appConfig.String("salt"))
	err = models.CreateUserAccount(username, enPassword)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/register", 302)
}
