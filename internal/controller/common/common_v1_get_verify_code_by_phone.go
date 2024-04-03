package common

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"gohub/internal/service"
	"gohub/utility/utils"

	"gohub/api/common/v1"
)

func (c *ControllerV1) GetVerifyCodeByPhone(ctx context.Context, req *v1.GetVerifyCodeByPhoneReq) (res *v1.GetVerifyCodeByPhoneRes, err error) {
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
	// 发送短信
	err = service.AliyunSms().Send(ctx, req.Mobile, verifyCode)
	if err != nil {
		g.Log().Errorf(ctx, "发送短信失败: %s\n", err.Error())
		err = gerror.New("服务器内部错误")
		return
	}
	// 发送成功，验证码保存到redis中
	err = g.Redis().SetEX(ctx, fmt.Sprintf("%s%s", cacheKeyPrefix, req.Mobile), verifyCode, expiredTime)
	if err != nil {
		g.Log().Errorf(ctx, "设置验证码缓存失败: %s\n", err.Error())
		err = gerror.New("服务器内部错误")
		return
	}
	return
}
