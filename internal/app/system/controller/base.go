package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	commonController "gohub/internal/app/common/controller"
)

type BaseController struct {
	commonController.BaseController
}

func (c *BaseController) Init(r *ghttp.Request) {
	c.BaseController.Init(r)
}
