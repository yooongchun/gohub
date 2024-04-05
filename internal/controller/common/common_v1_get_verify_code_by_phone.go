package common

import (
	"context"
	"gohub/api/common/v1"
	"gohub/internal/consts"
	"gohub/internal/service"
	"gohub/utility/errUtils"
)

func (c *ControllerV1) GetVerifyCodeByPhone(ctx context.Context, req *v1.GetVerifyCodeByPhoneReq) (res *v1.GetVerifyCodeByPhoneRes, err error) {
	//判断验证码是否正确
	verifyCaptcha(req.VerifyKey, req.VerifyCode)
	// 发送短信
	verifyCode := genVerifyCode()
	err = service.AliyunSms().Send(ctx, req.Mobile, verifyCode)
	errUtils.ErrIfNotNil(ctx, err, consts.InternalServerError)
	// 发送成功，验证码保存到redis中
	cacheVerifyCode(req.Mobile, verifyCode)
	return
}
