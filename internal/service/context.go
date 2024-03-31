// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gohub/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IContext interface {
		// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
		Init(r *ghttp.Request, customCtx *model.Context)
		// Get 获得上下文变量，如果没有设置，那么返回nil
		Get(ctx context.Context) *model.Context
		// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
		SetUser(ctx context.Context, ctxUser *model.ContextUser)
		// GetLoginUser 获取当前登陆用户信息
		GetLoginUser(ctx context.Context) *model.ContextUser
		// GetUserId 获取当前登录用户id
		GetUserId(ctx context.Context) uint64
	}
)

var (
	localContext IContext
)

func Context() IContext {
	if localContext == nil {
		panic("implement not found for interface IContext, forgot register?")
	}
	return localContext
}

func RegisterContext(i IContext) {
	localContext = i
}
