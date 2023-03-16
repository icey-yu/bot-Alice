package main

import (
	_ "bot-Alice/internal/core"
	_ "bot-Alice/internal/packed"
	"bot-Alice/internal/service"
)

func main() {
	//cmd.Main.Run(gctx.New())
	service.Login()

	// 停了
	select {}
}
