// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		MiddlewareCORS(r *ghttp.Request)
		// Ctx 自定义上下文对象
		Ctx(r *ghttp.Request)
		// AdminRequired 权限判断处理中间件
		AdminRequired(r *ghttp.Request)
		// LoginRequired 登录判断处理中间件
		LoginRequired(r *ghttp.Request)
		// ErrorHandler 错误处理中间件
		ErrorHandler(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
