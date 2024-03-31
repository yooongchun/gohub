// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOperateLog is the golang structure of table sys_operate_log for DAO operations like Where/Data.
type SysOperateLog struct {
	g.Meta          `orm:"table:sys_operate_log, do:true"`
	OperateId       interface{} // 日志主键
	Title           interface{} // 模块标题
	BusinessType    interface{} // 业务类型（0其它 1新增 2修改 3删除）
	Method          interface{} // 方法名称
	RequestMethod   interface{} // 请求方式
	OperatorType    interface{} // 操作类别（0其它 1后台用户 2手机端用户）
	OperateName     interface{} // 操作人员
	OperateUrl      interface{} // 请求URL
	OperateIp       interface{} // 主机地址
	OperateLocation interface{} // 操作地点
	OperateParam    interface{} // 请求参数
	ErrorMsg        interface{} // 错误消息
	OperateTime     *gtime.Time // 操作时间
}
