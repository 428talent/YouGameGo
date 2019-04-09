package security

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"yougame.com/yougame-server/util"
)

var AppSecret string

func init() {
	appConfig, err := config.NewConfig("ini", "./conf/app_local.conf")
	if err != nil {
		beego.Error(err)
	}
	AppSecret =  util.GetConfigItem("APPLICATION_SECRET","app_secret",appConfig,"")
}
