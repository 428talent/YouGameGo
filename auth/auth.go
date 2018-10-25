package auth

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"yougame.com/letauthsdk/client"
)

var AuthClient client.LetAuthClient

func init() {
	appConfig, err := config.NewConfig("ini", "./conf/app_local.conf")
	if err != nil {
		beego.Error(err)
	}
	baseUrl := appConfig.DefaultString("auth_server_address", "http://localhost:8080")
	AuthClient = client.BuildClient(baseUrl)
}
