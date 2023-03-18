// 初始化，包含隐式和显示

package core

import (
	"bot-Alice/internal/core/event"
	"bot-Alice/internal/core/login"
	"bot-Alice/internal/core/redis"
	_ "bot-Alice/internal/logic"
)

func Init() error {
	if err := login.Init(); err != nil {
		return err
	}

	redis.Init()
	event.Init()
	return nil
}
