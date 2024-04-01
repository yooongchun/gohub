package mail

import (
	"context"
	"gohub/internal/service"
	"gohub/utility/utils"
	"gopkg.in/gomail.v2"
)

func init() {
	service.RegisterMail(New())
}

type sMail struct {
}

func (s *sMail) Send(ctx context.Context, to, subject, html string) (err error) {
	sender := utils.GetConfig(ctx, "mail.qq.sender")
	host := utils.GetConfig(ctx, "mail.qq.host")
	token := utils.GetConfig(ctx, "mail.qq.token")
	port := utils.GetConfigInt(ctx, "mail.qq.port")

	m := gomail.NewMessage()
	//发送人
	m.SetHeader("From", sender)
	//接收人
	m.SetHeader("To", to)
	//主题
	m.SetHeader("Subject", subject)
	//内容
	m.SetBody("text/html", html)

	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer(host, port, sender, token)

	// 发送邮件
	err = d.DialAndSend(m)
	return
}

func New() *sMail {
	return &sMail{}
}
