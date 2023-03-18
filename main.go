package main

import (
	"bot-Alice/internal/core"
	_ "bot-Alice/internal/core"
	_ "bot-Alice/internal/packed"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()
	//cmd.Main.Run(gctx.New())
	if err := core.Init(); err != nil {
		g.Log().Errorf(ctx, "初始化失败:%+v", err)
		return
	} // 初始化

	// 停了
	select {}
}
