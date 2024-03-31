package v1

import "github.com/gogf/gf/v2/frame/g"

// EmptyRes 空响应结构体
type EmptyRes struct {
	g.Meta `mime:"application/json"`
}

// ListRes 列表响应结构体
type ListRes struct {
	CurrentPage int         `json:"current_page"`
	Total       interface{} `json:"total"`
}
