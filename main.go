package main

import (
	_ "bot-Alice/internal/core"
	_ "bot-Alice/internal/packed"
	"bot-Alice/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()
	//cmd.Main.Run(gctx.New())
	err := service.Login().Login()
	if err != nil {
		g.Log().Error(ctx, "登录失败", err)
	}

	// 停了
	select {}
}
