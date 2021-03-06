package transaction

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"time"
	"yougame.com/yougame-server/log"
	"yougame.com/yougame-server/util"
)

var TransactionServiceClient *TransactionClient

func init() {
	appConfig, err := config.NewConfig("ini", "./conf/app_local.conf")
	if err != nil {
		log.LogClient.Channel <- &log.Payload{
			Message: err.Error(),
			Level:   "error",
			Time:    time.Now(),
		}
	}
	mqAddress := util.GetConfigItem("APPLICATION_MQ_URL","rbq.address",appConfig,"amqp://localhost:5672/")
	client := TransactionClient{}
	err = client.Connect(ClientConfig{Address: mqAddress})
	if err != nil {
		beego.Error(err)
		log.LogClient.Channel <- &log.Payload{
			Message: err.Error(),
			Level:   "error",
			Time:    time.Now(),
		}
		return
	}
	log.LogClient.Channel <- &log.Payload{
		Message: "success connect to transaction service",
		Level:   "info",
		Time:    time.Now(),
	}
	if err != nil {
		beego.Error(err)
	}
	beego.Debug("连接账单服务成功")
	TransactionServiceClient = &client
}
