package sms

import (
	"context"
	"encoding/json"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"gohub/internal/service"
	"gohub/utility/utils"
)

func init() {
	service.RegisterAliyunSms(New())
}

func (s *sAliyunSms) Send(ctx context.Context, mobile string, data map[string]string) (err error) {
	var client *dysmsapi.Client
	client, err = dysmsapi.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(utils.GetConfig(ctx, "sms.aliyun.accessKeyId")),
		AccessKeySecret: tea.String(utils.GetConfig(ctx, "sms.aliyun.accessKeySecret"))})
	if err != nil {
		g.Log().Errorf(ctx, "初始化阿里云短信客户端失败: %s\n", err.Error())
		return
	}
	// 发送短信
	var templateParam []byte
	templateParam, err = json.Marshal(data)
	if err != nil {
		g.Log().Errorf(ctx, "序列化短信模板参数失败: %s\n", err.Error())
		return
	}
	request := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  tea.String(mobile),
		SignName:      tea.String(utils.GetConfig(ctx, "sms.aliyun.signName")),
		TemplateCode:  tea.String(utils.GetConfig(ctx, "sms.aliyun.templateCode")),
		TemplateParam: tea.String(string(templateParam)),
	}
	var resp *dysmsapi.SendSmsResponse
	resp, err = client.SendSms(request)
	if err != nil {
		g.Log().Errorf(ctx, "发送短信失败: %s\n", err.Error())
		return
	}
	if *resp.Body.Code != "OK" {
		g.Log().Errorf(ctx, "发送短信失败: %s\n", *resp.Body.Message)
		err = gerror.New("响应错误")
		return
	}
	return
}

func New() *sAliyunSms {
	return &sAliyunSms{}
}

type sAliyunSms struct {
}
