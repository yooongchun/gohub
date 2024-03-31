package service

import (
	"context"
	"gohub/api/v1/system"
	"gohub/internal/app/system/model"
)

type ISysLoginLog interface {
	Invoke(ctx context.Context, data *model.LoginLogParams)
	List(ctx context.Context, req *system.LoginLogSearchReq) (res *system.LoginLogSearchRes, err error)
	DeleteLoginLogByIds(ctx context.Context, ids []int) (err error)
	ClearLoginLog(ctx context.Context) (err error)
}

var localSysLoginLog ISysLoginLog

func SysLoginLog() ISysLoginLog {
	if localSysLoginLog == nil {
		panic("implement not found for interface ISysLoginLog, forgot register?")
	}
	return localSysLoginLog
}

func RegisterSysLoginLog(i ISysLoginLog) {
	localSysLoginLog = i
}
