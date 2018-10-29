package web

import (
	"github.com/astaxie/beego"
	"yougame.com/letauthsdk/auth"
	"yougame.com/yougame-server/models"
)

func SetPageAuthInfo(c beego.Controller, claims *auth.UserClaims) {
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
	} else {
		c.Data["isLogin"] = false
	}
}
