package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "gohub/api/v1/common"
	"gohub/internal/app/system/model/entity"
)

// LoginLogSearchReq 查询列表请求参数
type LoginLogSearchReq struct {
	g.Meta        `path:"/log/login/list" tags:"登录日志管理" method:"get" summary:"日志列表"`
	LoginName     string `p:"userName"`      //登陆名
	Status        string `p:"status"`        //状态
	Ipaddr        string `p:"ipaddr"`        //登录地址
	SortName      string `p:"orderByColumn"` //排序字段
	SortOrder     string `p:"isAsc"`         //排序方式
	LoginLocation string `p:"loginLocation"` //登录地点
	commonApi.PageReq
}

type LoginLogSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*entity.SysLoginLog `json:"list"`
}

type LoginLogDelReq struct {
	g.Meta `path:"/log/login/delete" tags:"登录日志管理" method:"delete" summary:"删除日志"`
	Ids    []int `p:"ids" v:"required#ids必须"`
}

type LoginLogDelRes struct {
}

type LoginLogClearReq struct {
	g.Meta `path:"/log/login/clear" tags:"登录日志管理" method:"delete" summary:"清除日志"`
}

type LoginLogClearRes struct {
}
