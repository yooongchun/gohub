package sms

import (
	"context"
	"encoding/json"
	aliyunsmsclient "github.com/KenmyZhang/aliyun-communicate"
	"github.com/gogf/gf/v2/frame/g"
	"gohub/internal/service"
	"gohub/utility/errUtils"
)

func init() {
	service.RegisterAliyunSms(New())
}

func New() *sAliyunSms {
	return &sAliyunSms{}
}

type sAliyunSms struct {
}

func (s *sAliyunSms) Send(ctx context.Context, mobile string, data map[string]string) bool {
	smsClient := aliyunsmsclient.New(g.Cfg().MustGet(ctx, "sms.aliyun.endpoint").String())
	templateParam, err := json.Marshal(data)
	errUtils.ErrIfNotNil(ctx, err)
	result, err := smsClient.Execute(
		g.Cfg().MustGet(ctx, "sms.aliyun.accessKeyId").String(),
		g.Cfg().MustGet(ctx, "sms.aliyun.accessKeySecret").String(),
		mobile,
		g.Cfg().MustGet(ctx, "sms.aliyun.signName").String(),
		g.Cfg().MustGet(ctx, "sms.aliyun.templateCode").String(),
		string(templateParam),
	)
	errUtils.ErrIfNotNil(ctx, err)
	return result.IsSuccessful()
}
