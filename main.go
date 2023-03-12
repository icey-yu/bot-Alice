package main

import (
	_ "bot-Alice/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"bot-Alice/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
