package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "gohub/api/v1/common"
	"gohub/internal/app/system/model"
)

// SysOperateLogSearchReq 分页请求参数
type SysOperateLogSearchReq struct {
	g.Meta        `path:"/log/operate/list" tags:"操作日志" method:"get" summary:"操作日志列表"`
	Title         string `p:"title"`         //系统模块
	RequestMethod string `p:"requestMethod"` //请求方式
	OperateName   string `p:"operateName"`   //操作人员
	commonApi.PageReq
	commonApi.Author
}

// SysOperateLogSearchRes 列表返回结果
type SysOperateLogSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*model.SysOperateLogListRes `json:"list"`
}

// SysOperateLogGetReq 获取一条数据请求
type SysOperateLogGetReq struct {
	g.Meta `path:"/log/operate/get" tags:"操作日志" method:"get" summary:"获取操作日志信息"`
	commonApi.Author
	OperateId uint64 `p:"operateId" v:"required#主键必须"` //通过主键获取
}

// SysOperateLogGetRes 获取一条数据结果
type SysOperateLogGetRes struct {
	g.Meta `mime:"application/json"`
	*model.SysOperateLogInfoRes
}

// SysOperateLogDeleteReq 删除数据请求
type SysOperateLogDeleteReq struct {
	g.Meta `path:"/log/operate/delete" tags:"操作日志" method:"delete" summary:"删除操作日志"`
	commonApi.Author
	OperateIds []uint64 `p:"operateIds" v:"required#主键必须"` //通过主键删除
}

// SysOperateLogDeleteRes 删除数据返回
type SysOperateLogDeleteRes struct {
	commonApi.EmptyRes
}

type SysOperateLogClearReq struct {
	g.Meta `path:"/log/operate/clear" tags:"操作日志" method:"delete" summary:"清除日志"`
	commonApi.Author
}

type SysOperateLogClearRes struct {
	commonApi.EmptyRes
}
