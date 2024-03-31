package controller

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
	"gohub/api/v1/system"
	commonService "gohub/internal/app/common/service"
	"gohub/internal/app/system/model"
	"gohub/internal/app/system/service"
	"gohub/utility/utils"
)

var Login = loginController{}

type loginController struct {
	BaseController
}

// Login 用户登陆
func (c *loginController) Login(ctx context.Context, req *system.UserLoginReq) (res *system.UserLoginRes, err error) {
	var (
		user  *model.LoginUserRes
		token string
	)
	// 判断验证码是否正确
	debug := gmode.IsDevelop()
	if !debug {
		if !commonService.Captcha().VerifyCaptcha(req.VerifyKey, req.VerifyCode) {
			err = gerror.New("验证码错误")
			return
		}
	}
	ip := utils.GetClientIp(ctx)
	userAgent := utils.GetUserAgent(ctx)
	user, err = service.SysUser().GetAdminUserByUsernamePassword(ctx, req)
	if err != nil {
		// 保存登陆失败的日志信息
		service.SysLoginLog().Invoke(ctx, &model.LoginLogParams{
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
	// 保存登陆成功的日志信息
	service.SysLoginLog().Invoke(ctx, &model.LoginLogParams{
		Status:    1,
		Username:  req.Username,
		Ip:        ip,
		UserAgent: userAgent,
		Msg:       "登陆成功",
		Module:    "系统后台",
	})
	key := gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.USerPassword)
	if g.Cfg().MustGet(ctx, "ghToken.multiLogin").Bool() {
		key = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.USerPassword+ip+userAgent)
	}
	user.USerPassword = ""
	token, err = service.GhToken().GenerateToken(ctx, key, user)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("登陆失败，后端服务异常")
		return
	}
	res = &system.UserLoginRes{
		UserInfo: user,
		Token:    token,
	}
	return
}

// Logout 用户登出
func (c *loginController) Logout(ctx context.Context, req *system.UserLogoutReq) (res *system.UserLogoutRes, err error) {
	err = service.GhToken().RemoveToken(ctx, service.GhToken().GetRequestToken(g.RequestFromCtx(ctx)))
	return
}
