package mail

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/jordan-wright/email"
	"net/smtp"
	"yougame.com/yougame-server/util"
)

var mailClient *email.Pool

func init() {
	appConfig, err := config.NewConfig("ini", "./conf/app_local.conf")
	if err != nil {
		beego.Error(err)
	}

	username := util.GetConfigItem("APPLICATION_MAIL_USERNAME","mail.username",appConfig,"")
	password := util.GetConfigItem("APPLICATION_MAIL_PASSWORD","mail.password",appConfig,"")
	host := util.GetConfigItem("APPLICATION_MAIL_HOST","mail.host",appConfig,"")
	address := util.GetConfigItem("APPLICATION_MAIL_ADDRESS","mail.address",appConfig,"")

	p, err := email.NewPool(
		address,
		4,
		smtp.PlainAuth("", username, password, host),
	)
	if err != nil {
		beego.Error(err)
	}
	mailClient = p
}

func SendMail(email *email.Email) {
	//go func() {
	//	err := mailClient.Send(email, 10*time.Second)
	//	if err != nil {
	//		beego.Error(err)
	//	}
	//}()
}
