package log

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"yougame.com/yougame-server/util"
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
	applicationId := util.GetConfigItem("APPLICATION_LOG_ID","log.application.id",appConfig,"repository")
	instanceId := util.GetConfigItem("APPLICATION_LOG_INSTANCE_ID","log.application.instance",appConfig,"main-server")
	address := util.GetConfigItem("APPLICATION_LOG_HOST","log.application.address",appConfig,"http://localhost:6000")
	LogClient = NewClient(applicationId, instanceId, address)
}
