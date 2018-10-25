package web

import (
	"github.com/astaxie/beego"
	"yougame.com/letauthsdk/auth"
	AppAuth "yougame.com/yougame-server/auth"
)

func SetPageAuthInfo(c beego.Controller, claims *auth.UserClaims) {
	if claims != nil {
		user, err := AppAuth.AuthClient.GetUser(claims.UserId)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Nickname"] = user.Profile.Nickname
		if len(user.Profile.Avatar) == 0 {
			c.Data["Avatar"] = "/static/img/user.png"
		} else {
			c.Data["Avatar"] = AppAuth.AuthClient.BaseUrl + "/" + user.Profile.Avatar
		}
		c.Data["isLogin"] = true
	} else {
		c.Data["isLogin"] = false
	}
}
