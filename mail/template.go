package mail

import (
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/matcornic/hermes"
	"net/textproto"
	"strconv"
	"time"
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
			Logo:      "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
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
func RenderVerifyCodeMail(user *models.User, code int) *hermes.Email {
	return &hermes.Email{
		Body: hermes.Body{
			Name: user.Username,
			Intros: []string{
				"您正在使用YouGame账户服务，以下是您的验证码",
			},
			Dictionary: []hermes.Entry{
				{Key: "验证码", Value: strconv.Itoa(code)},
			},
			Outros: []string{
				"请勿将验证码泄露给任何人，避免给您的账号安全带来不必要的麻烦。",
			},
		},
	}

}

type ReceiptMailGoods struct {
	GameName string
	GoodName string
	Price    float64
}

type ReceiptMailModel struct {
	Name            string
	Items           []*ReceiptMailGoods
	OrderId         int64
	TransactionTime *time.Time
	TotalPrice      float64
	TransactionId   int
}

func SendReceiptMail(model ReceiptMailModel, to string) error {
	template := RenderReceiptMail(model)
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
		Subject: "感谢您的购买",
		Text:    []byte(textString),
		HTML:    []byte(htmlString),
		Headers: textproto.MIMEHeader{},
	}
	SendMail(mail)
	return nil
}
func RenderReceiptMail(model ReceiptMailModel) *hermes.Email {
	tableData := make([][]hermes.Entry, 0)
	for _, good := range model.Items {
		tableData = append(tableData, []hermes.Entry{
			{Key: "游戏", Value: good.GameName},
			{Key: "商品", Value: good.GoodName},
			{Key: "金额", Value: fmt.Sprintf("¥%.2f", good.Price)},
		})
	}
	return &hermes.Email{
		Body: hermes.Body{
			Name: model.Name,
			Intros: []string{
				"感谢您在YouGame的购物，已将购买的商品添加至您的仓库中",
			},
			Dictionary: []hermes.Entry{
				{Key: "订单号", Value: strconv.Itoa(int(model.OrderId))},
				{Key: "交易号", Value: strconv.Itoa(int(model.TransactionId))},
				{Key: "交易时间", Value: model.TransactionTime.Format("2006-01-02 15:04:05")},
				{Key: "总计", Value: fmt.Sprintf("¥%.2f", model.TotalPrice)},
			},
			Outros: []string{
				"如果您对购买内容有疑问，请联系服务人员。(请勿回复本邮件)",
			},
			Table: hermes.Table{
				Data: tableData,
			},
			Greeting: "您好",
			Signature: "",

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

func SendVerifyCodeEmail(user *models.User, to string, code int) error {
	template := RenderVerifyCodeMail(user, code)
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
		Subject: "YouGame验证码",
		Text:    []byte(textString),
		HTML:    []byte(htmlString),
		Headers: textproto.MIMEHeader{},
	}
	SendMail(mail)
	return nil
}
