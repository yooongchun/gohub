package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"gohub/api/common/v1"
	"gohub/internal/model/entity"
)

// GetLoginLogListReq 查询列表请求参数
type GetLoginLogListReq struct {
	g.Meta        `path:"/login-logs" tags:"日志管理" method:"get" summary:"登陆日志列表"`
	LoginName     string `p:"userName"`      //登陆名
	Status        string `p:"status"`        //状态
	Ipaddr        string `p:"ipaddr"`        //登录地址
	SortName      string `p:"orderByColumn"` //排序字段
	SortOrder     string `p:"isAsc"`         //排序方式
	LoginLocation string `p:"loginLocation"` //登录地点
	v1.PageReq
}

type GetLoginLogListRes struct {
	g.Meta `mime:"application/json"`
	v1.ListRes
	List []*entity.SysLoginLog `json:"list"`
}

type DeleteLoginLogListReq struct {
	g.Meta `path:"/login-logs" tags:"日志管理" method:"delete" summary:"删除登陆日志"`
	Ids    []int `p:"ids" v:"required#ids必须"`
}

type DeleteLoginLogListRes struct {
	v1.EmptyRes
}

type ClearLoginLogReq struct {
	g.Meta `path:"/login-logs/clear" tags:"日志管理" method:"delete" summary:"清除登陆日志"`
}

type ClearLoginLogRes struct {
	v1.EmptyRes
}
