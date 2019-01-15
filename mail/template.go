package mail

import (
	"github.com/jordan-wright/email"
	"github.com/matcornic/hermes"
	"net/textproto"
	"yougame.com/yougame-server/models"
)

var hermesApp hermes.Hermes

func init() {
	h := hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: "You Game",
			Link: "https://example-hermes.com/",
			// Optional product logo
			Logo: "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
			Copyright: "Copyright © 2019 YouGame Project. All rights reserved.",
		},
	}
	hermesApp = h
}

func RenderWelcomeMail(user *models.User) *hermes.Email {
	return &hermes.Email{
		Body: hermes.Body{
			Name: user.Username,
			Intros: []string{
				"欢迎使用You Game服务，您已经成功的创建了账户。",
			},
		},
	}

}

func SendWelcomeEmail(user *models.User, to string) error {
	template := RenderWelcomeMail(user)
	htmlString, err := hermesApp.GenerateHTML(*template)
	if err != nil {
		return err
	}
	textString, err := hermesApp.GeneratePlainText(*template)
	if err != nil {
		return err
	}
	mail := &email.Email{
		To:      []string{to,},
		From:    "YouGame Project <takayamaaren@sina.com>",
		Subject: "注册成功",
		Text:    []byte(textString),
		HTML:    []byte(htmlString),
		Headers: textproto.MIMEHeader{},
	}
	SendMail(mail)
	return nil
}
