package system

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
	"gohub/internal/model"
	"gohub/internal/service"
	"gohub/utility/utils"

	"github.com/gogf/gf/v2/errors/gerror"

	"gohub/api/system/v1"
)

func (c *ControllerV1) UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	var (
		user  *model.LoginUserRes
		token string
	)
	//判断验证码是否正确
	debug := gmode.IsDevelop()
	if !debug {
		if !service.Captcha().VerifyCaptcha(req.VerifyKey, req.VerifyCode) {
			err = gerror.New("验证码输入错误")
			return
		}
	}
	ip := utils.GetClientIp(ctx)
	userAgent := utils.GetUserAgent(ctx)
	user, err = service.SysUser().GetAdminUserByUsernamePassword(ctx, req)
	if err != nil {
		// 保存登录失败的日志信息
		service.SysLoginLog().Invoke(gctx.New(), &model.LoginLogParams{
			Status:    0,
			Username:  req.Username,
			Ip:        ip,
			UserAgent: userAgent,
			Msg:       err.Error(),
			Module:    "系统后台",
		})
		return
	}
	err = service.SysUser().UpdateLoginInfo(ctx, user.Id, ip)
	if err != nil {
		return
	}
	// 报存登录成功的日志信息
	service.SysLoginLog().Invoke(gctx.New(), &model.LoginLogParams{
		Status:    1,
		Username:  req.Username,
		Ip:        ip,
		UserAgent: userAgent,
		Msg:       "登录成功",
		Module:    "系统后台",
	})
	key := gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword)
	if g.Cfg().MustGet(ctx, "sSysLoginLog.multiLogin").Bool() {
		key = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword+ip+userAgent)
	}
	user.UserPassword = ""
	token, err = service.GhToken().GenerateToken(ctx, key, user)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("登录失败，后端服务出现错误")
		return
	}
	//获取用户菜单数据
	res = &v1.UserLoginRes{
		UserInfo: user,
		Token:    token,
	}
	return
}
