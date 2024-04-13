package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"gohub/api/common/v1"
	"gohub/internal/model/entity"
)

// GetOperateLogListReq 分页请求参数
type GetOperateLogListReq struct {
	g.Meta        `path:"/operate-logs" tags:"操作日志" method:"get" summary:"操作日志列表"`
	Title         string `p:"title"`         //系统模块
	RequestMethod string `p:"requestMethod"` //请求方式
	OperateName   string `p:"operateName"`   //操作人员
	v1.PageReq
	v1.Author
}

// GetOperateLogListRes 列表返回结果
type GetOperateLogListRes struct {
	g.Meta `mime:"application/json"`
	v1.ListRes
	List []*entity.SysLoginLog `json:"list"`
}

// GetOperateLogOneReq 获取一条数据请求
type GetOperateLogOneReq struct {
	g.Meta `path:"/operate-log/:operateId" tags:"操作日志" method:"get" summary:"获取操作日志信息"`
	v1.Author
}

// GetOperateLogOneRes 获取一条数据结果
type GetOperateLogOneRes struct {
	g.Meta `mime:"application/json"`
	*entity.SysOperateLog
}

// DeleteOperateLogListReq 删除数据请求
type DeleteOperateLogListReq struct {
	g.Meta `path:"/operate-logs" tags:"操作日志" method:"delete" summary:"删除操作日志"`
	v1.Author
	OperateIds []uint64 `p:"operateIds" v:"required#主键必须"` //通过主键删除
}

// DeleteOperateLogListRes 删除数据结果
type DeleteOperateLogListRes struct {
	v1.EmptyRes
}

// ClearOperateLogReq 清除操作日志请求
type ClearOperateLogReq struct {
	g.Meta `path:"/operate-logs/clear" tags:"操作日志" method:"delete" summary:"清除日志"`
	v1.Author
}

// ClearOperateLogRes 清除操作日志结果
type ClearOperateLogRes struct {
	v1.EmptyRes
}
