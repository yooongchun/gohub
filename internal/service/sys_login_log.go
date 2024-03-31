// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "gohub/api/system/v1"
	"gohub/internal/model"
)

type (
	ISysLoginLog interface {
		// Invoke 异步写入登录日志
		Invoke(ctx context.Context, data *model.LoginLogParams)
		// List 获取登录日志列表
		List(ctx context.Context, req *v1.GetLoginLogListReq) (res *v1.GetLoginLogListRes, err error)
		// DeleteLoginLogByIds 删除登录日志
		DeleteLoginLogByIds(ctx context.Context, ids []int) (err error)
		// ClearLoginLog 清空登录日志
		ClearLoginLog(ctx context.Context) (err error)
	}
)

var (
	localSysLoginLog ISysLoginLog
)

func SysLoginLog() ISysLoginLog {
	if localSysLoginLog == nil {
		panic("implement not found for interface ISysLoginLog, forgot register?")
	}
	return localSysLoginLog
}

func RegisterSysLoginLog(i ISysLoginLog) {
	localSysLoginLog = i
}
