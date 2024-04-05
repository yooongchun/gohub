package common

import (
	"bytes"
	"context"
	"gohub/api/common/v1"
	"gohub/internal/consts"
	"gohub/internal/service"
	"gohub/utility/errUtils"
	"html/template"
	"strconv"
)

func (c *ControllerV1) GetVerifyCodeByEmail(ctx context.Context, req *v1.GetVerifyCodeByEmailReq) (res *v1.GetVerifyCodeByEmailRes, err error) {
	//判断验证码是否正确
	verifyCaptcha(req.VerifyKey, req.VerifyCode)
	//生成验证码
	verifyCode := genVerifyCode()
	// 渲染邮件模板
	tmpl := template.Must(template.ParseFiles(tmplFilePath))
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, map[string]string{"VerifyCode": verifyCode, "ExpiredTime": strconv.FormatInt(expiredTime/60, 10)})
	errUtils.ErrIfNotNil(ctx, err, consts.InternalServerError)
	// 发送邮件
	err = service.Mail().Send(ctx, req.Email, "永春小站验证码", buf.String())
	errUtils.ErrIfNotNil(ctx, err, consts.InternalServerError)
	// 发送成功，验证码保存到redis中
	cacheVerifyCode(req.Email, verifyCode)
	return
}
