package common

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"gohub/api/common/v1"
	"gohub/internal/service"
	"gohub/utility/utils"
	"html/template"
	"strconv"
)

func (c *ControllerV1) GetVerifyCodeByEmail(ctx context.Context, req *v1.GetVerifyCodeByEmailReq) (res *v1.GetVerifyCodeByEmailRes, err error) {
	//判断验证码是否正确
	if !service.Captcha().VerifyCaptcha(req.VerifyKey, req.VerifyCode) {
		g.Log().Errorf(ctx, "验证码输入错误: verifyKey=%s,verifyCode=%s\n", req.VerifyKey, req.VerifyCode)
		err = gerror.New("验证码输入错误")
		return
	}
	// 配置项
	cacheKeyPrefix := utils.GetConfig(ctx, "verifyCode.cacheKeyPrefix")
	expiredTime := int64(utils.GetConfigInt(ctx, "verifyCode.expiredTime"))

	//生成验证码
	verifyCode := grand.Digits(6)
	// 渲染邮件模板
	tmpl := template.Must(template.ParseFiles("resource/template/email.html"))
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, map[string]string{"VerifyCode": verifyCode, "ExpiredTime": strconv.FormatInt(expiredTime/60, 10)})
	if err != nil {
		g.Log().Errorf(ctx, "模板渲染失败: %s\n", err.Error())
		err = gerror.New("服务器内部错误")
		return
	}
	// 发送邮件
	err = service.Mail().Send(ctx, req.Email, "永春小站验证码", buf.String())
	if err != nil {
		g.Log().Errorf(ctx, "发送邮件失败: %s\n", err.Error())
		err = gerror.New("服务器内部错误")
		return
	}
	// 发送成功，验证码保存到redis中
	err = g.Redis().SetEX(ctx, fmt.Sprintf("%s%s", cacheKeyPrefix, req.Email), verifyCode, expiredTime)
	if err != nil {
		g.Log().Errorf(ctx, "设置验证码缓存失败: %s\n", err.Error())
		err = gerror.New("服务器内部错误")
		return
	}
	return
}
