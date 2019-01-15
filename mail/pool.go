package mail

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/jordan-wright/email"
	"net/smtp"
	"time"
)

var mailClient *email.Pool

func init() {
	appConfig, err := config.NewConfig("ini", "./conf/app_local.conf")
	if err != nil {
		beego.Error(err)
	}

	p, err := email.NewPool(
		appConfig.String("mail.address"),
		4,
		smtp.PlainAuth("", appConfig.String("mail.username"), appConfig.String("mail.password"), appConfig.String("mail.host")),
	)
	if err != nil {
		beego.Error(err)
	}
	mailClient = p
}

func SendMail(email *email.Email) {
	go func() {
		err := mailClient.Send(email, 10*time.Second)
		if err != nil {
			beego.Error(err)
		}
	}()
}
