package mail

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	emailClient "github.com/jordan-wright/email"
	"gohub/internal/model"
	"gohub/internal/service"
	"gohub/utility/errUtils"
	"net/smtp"
)

func init() {
	service.RegisterMail(New())
}

type sMail struct {
}

func (s *sMail) Send(ctx context.Context, email model.Email) bool {
	e := emailClient.NewEmail()
	e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.HTML

	err := e.Send(
		fmt.Sprintf("%v:%v", g.Cfg().MustGet(ctx, "mail.qq.host").String(), g.Cfg().MustGet(ctx, "mail.qq.port").String()),
		smtp.PlainAuth(
			"",
			g.Cfg().MustGet(ctx, "mail.qq.account").String(),
			g.Cfg().MustGet(ctx, "mail.qq.password").String(),
			g.Cfg().MustGet(ctx, "mail.qq.host").String(),
		),
	)
	errUtils.ErrIfNotNil(ctx, err)
	return true
}

func New() *sMail {
	return &sMail{}
}
