package log

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

var LogClient *Client

func init() {
	initLogClient()
}
func initLogClient() {
	appConfig, err := config.NewConfig("ini", "./conf/app_local.conf")
	if err != nil {
		beego.Error(err)
	}
	applicationId := appConfig.DefaultString("log.application.id", "repository")
	instanceId := appConfig.DefaultString("log.application.instance", "main-server")
	address := appConfig.DefaultString("log.application.address", "http://localhost:5002")
	LogClient = NewClient(applicationId, instanceId, address)
}
