// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "gohub/api/system/v1"
	"gohub/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IOperateLog interface {
		// OperationLog 操作日志写入
		OperationLog(r *ghttp.Request)
		Invoke(ctx context.Context, data *model.SysOperateLogAdd)
		// List 操作日志列表
		List(ctx context.Context, req *v1.GetOperateLogListReq) (listRes *v1.GetOperateLogListRes, err error)
		// GetByOperateId 根据id获取操作日志
		GetByOperateId(ctx context.Context, operateId uint64) (res *model.SysOperateLogInfoRes, err error)
		// DeleteByIds 根据id删除操作日志
		DeleteByIds(ctx context.Context, ids []uint64) (err error)
		// ClearLog 清除操作日志
		ClearLog(ctx context.Context) (err error)
	}
)

var (
	localOperateLog IOperateLog
)

func OperateLog() IOperateLog {
	if localOperateLog == nil {
		panic("implement not found for interface IOperateLog, forgot register?")
	}
	return localOperateLog
}

func RegisterOperateLog(i IOperateLog) {
	localOperateLog = i
}
