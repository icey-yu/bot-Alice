package event

import (
	"bot-Alice/internal/global"
	"bot-Alice/internal/service"
)

// Init 事件初始化
func Init() {
	global.Alice.GroupMessageEvent.Subscribe(service.GroupMessageEvent().Event)
}
