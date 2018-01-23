package api_web

import (
	"github.com/astaxie/beego"
	"github.com/pborman/uuid"
	"time"
	"you_game_go/models"
)

type UserLoginController struct {
	beego.Controller
}

func (c *UserLoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	user := models.User{
		Username: username,
		Password: password,
	}
	isValidate := models.CheckUserValidate(&user)
	if isValidate {
		beego.Info(user)
		loginToken := models.Token{
			UserId:    user.Id,
			LoginTime: time.Now(),
		}
		userToken := uuid.New()
		models.InsertTokenToRedis(&loginToken, userToken)
		c.Ctx.SetCookie("token", userToken)
		c.Redirect("/",302)
	} else {
		beego.Error("用户名或密码错误")
	}
	c.Redirect("/login", 302)
}
