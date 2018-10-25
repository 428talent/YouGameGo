package security

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)
var AppSecret string
func init() {
	appConfig, err := config.NewConfig("ini", "./conf/app_local.conf")
	if err != nil {
		beego.Error(err)
	}
	AppSecret = appConfig.String("app_secret")
}