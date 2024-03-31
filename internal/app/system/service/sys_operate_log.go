package service

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"gohub/api/v1/system"
	"gohub/internal/app/system/model"
)

type IOperateLog interface {
	OperationLog(r *ghttp.Request)
	Invoke(ctx context.Context, data *model.SysOperateLogAdd)
	List(ctx context.Context, req *system.SysOperateLogSearchReq) (listRes *system.SysOperateLogSearchRes, err error)
	GetByOperateId(ctx context.Context, operateId uint64) (res *model.SysOperateLogInfoRes, err error)
	DeleteByIds(ctx context.Context, ids []uint64) (err error)
	ClearLog(ctx context.Context) (err error)
}

var localOperateLog IOperateLog

func OperateLog() IOperateLog {
	if localOperateLog == nil {
		panic("implement not found for interface IOperateLog, forgot register?")
	}
	return localOperateLog
}

func RegisterOperateLog(i IOperateLog) {
	localOperateLog = i
}
