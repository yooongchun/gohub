// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOperateLog is the golang structure for table sys_operate_log.
type SysOperateLog struct {
	OperateId       uint64      `json:"operateId"       description:"日志主键"`
	Title           string      `json:"title"           description:"模块标题"`
	BusinessType    int         `json:"businessType"    description:"业务类型（0其它 1新增 2修改 3删除）"`
	Method          string      `json:"method"          description:"方法名称"`
	RequestMethod   string      `json:"requestMethod"   description:"请求方式"`
	OperatorType    int         `json:"operatorType"    description:"操作类别（0其它 1后台用户 2手机端用户）"`
	OperateName     string      `json:"operateName"     description:"操作人员"`
	OperateUrl      string      `json:"operateUrl"      description:"请求URL"`
	OperateIp       string      `json:"operateIp"       description:"主机地址"`
	OperateLocation string      `json:"operateLocation" description:"操作地点"`
	OperateParam    string      `json:"operateParam"    description:"请求参数"`
	ErrorMsg        string      `json:"errorMsg"        description:"错误消息"`
	OperateTime     *gtime.Time `json:"operateTime"     description:"操作时间"`
}
