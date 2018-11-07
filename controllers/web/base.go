package web

import (
	"github.com/astaxie/beego"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
)

type WebController struct {
	beego.Controller
	User models.User
}

func (c *WebController) SetPageAuthInfo(claims *security.UserClaims) {
	if claims != nil {
		user, err := models.GetUserById(claims.UserId)
		if err != nil {
			c.Data["isLogin"] = false
			return
		}
		err = user.ReadProfile()
		if err != nil {
			c.Data["isLogin"] = false
			return
		}
		c.Data["Nickname"] = user.Profile.Nickname
		if len(user.Profile.Avatar) == 0 {
			c.Data["Avatar"] = "/static/img/user.png"
		} else {
			c.Data["Avatar"] = user.Profile.Avatar
		}
		c.Data["isLogin"] = true
		c.User = *user
	} else {
		c.Data["isLogin"] = false
	}
}

func (c *WebController) LoadRequestUser(claims *security.UserClaims) {
	if claims != nil {
		user, err := models.GetUserById(claims.UserId)
		if err != nil {
			return
		}
		c.User = *user
	}
}
