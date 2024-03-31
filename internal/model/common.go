package model

// PageReq 分页请求公共参数
type PageReq struct {
	DateRange []string `p:"dateRange"` //日起范围
	PageNum   int      `p:"pageNum"`   //页码
	PageSize  int      `p:"pageSize"`  //每页数
	OrderBy   string   //排序字段
}

// ListRes 分页返回公共参数
type ListRes struct {
	CurrentPage int         `json:"currentPage"` //当前页
	Total       interface{} `json:"total"`       //总数
}
