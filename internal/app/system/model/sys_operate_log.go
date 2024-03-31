package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
	"net/url"
)

// SysOperateLogAdd 添加操作日志参数
type SysOperateLogAdd struct {
	User         *ContextUser
	Url          *url.URL
	Params       g.Map
	Method       string
	ClientIp     string
	OperatorType int
}

// SysOperateLogInfoRes is the golang structure for table sys_operate_log.
type SysOperateLogInfoRes struct {
	gmeta.Meta      `orm:"table:sys_operate_log"`
	OperateId       uint64                      `orm:"operate_id,primary" json:"operateId"` // 日志编号
	Title           string                      `orm:"title" json:"title"`                  // 系统模块
	BusinessType    int                         `orm:"business_type" json:"businessType"`   // 操作类型
	Method          string                      `orm:"method" json:"method"`                // 操作方法
	RequestMethod   string                      `orm:"request_method" json:"requestMethod"` // 请求方式
	OperatorType    int                         `orm:"operator_type" json:"operatorType"`   // 操作类别
	OperateName     string                      `orm:"operate_name" json:"operateName"`     // 操作人员
	DeptName        string                      `orm:"dept_name" json:"deptName"`           // 部门名称
	LinkedDeptName  *LinkedSysOperateLogSysDept `orm:"with:dept_id=dept_name" json:"linkedDeptName"`
	OperateUrl      string                      `orm:"operate_url" json:"operateUrl"`           // 请求URL
	OperateIp       string                      `orm:"operate_ip" json:"operateIp"`             // 主机地址
	OperateLocation string                      `orm:"operate_location" json:"operateLocation"` // 操作地点
	OperateParam    string                      `orm:"operate_param" json:"operateParam"`       // 请求参数
	ErrorMsg        string                      `orm:"error_msg" json:"errorMsg"`               // 错误消息
	OperateTime     *gtime.Time                 `orm:"operate_time" json:"operateTime"`         // 操作时间
}

type LinkedSysOperateLogSysDept struct {
	gmeta.Meta `orm:"table:sys_dept"`
	DeptId     int64  `orm:"dept_id" json:"deptId"`     // 部门id
	DeptName   string `orm:"dept_name" json:"deptName"` // 部门名称
}

type SysOperateLogListRes struct {
	OperateId       uint64                      `json:"operated"`
	Title           string                      `json:"title"`
	RequestMethod   string                      `json:"requestMethod"`
	OperateName     string                      `json:"operateName"`
	DeptName        string                      `json:"deptName"`
	LinkedDeptName  *LinkedSysOperateLogSysDept `orm:"with:dept_id=dept_name" json:"linkedDeptName"`
	OperateUrl      string                      `json:"operateUrl"`
	OperateIp       string                      `json:"operateIp"`
	OperateLocation string                      `json:"operateLocation"`
	OperateParam    string                      `json:"operateParam"`
	OperateTime     *gtime.Time                 `json:"operateTime"`
}
