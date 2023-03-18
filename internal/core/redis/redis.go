package redis

import (
	"bot-Alice/internal/global"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/redis/go-redis/v9"
)

// Init 初始化redis
func Init() {
	ctx := gctx.New()
	host := g.Cfg().MustGet(ctx, "redis.host").String()
	port := g.Cfg().MustGet(ctx, "redis.port").String()
	addr := host + ":" + port
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: g.Cfg().MustGet(ctx, "redis.psw").String(), // no password set
		DB:       g.Cfg().MustGet(ctx, "redis.db").Int(),
	})

	println(global.Redis.Ping(ctx).String())
}
