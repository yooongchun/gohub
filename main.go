package main

import (
	"github.com/gogf/gf/v2/frame/g"
	_ "gohub/internal/logic"
	"gohub/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"gohub/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	g.Log().Info(gctx.New(), "send email...")
	service.Mail().Send(gctx.New(), "1729465178@qq.com", "Test Gohub", "<h1>Hello!!!</h1>")
	cmd.Main.Run(gctx.GetInitCtx())
}
