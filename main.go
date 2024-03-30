package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	_ "gohub/internal/app/boot"
	_ "gohub/internal/packed"

	"gohub/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
