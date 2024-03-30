package main

import (
	_ "gohub/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gohub/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
