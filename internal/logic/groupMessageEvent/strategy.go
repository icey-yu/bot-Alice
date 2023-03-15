package groupMessageEvent

import (
	"bot-Alice/internal/service"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
)

type (
	sGroupMessageEvent struct{}

	// groupMessageEventPerformer 具体的事件执行
	groupMessageEventPerformer struct {
		Client          *client.QQClient
		Event           *message.GroupMessage
		strategyHandler []func() bool
	}
)

func new_() *sGroupMessageEvent {
	return &sGroupMessageEvent{}
}

func newPerformer(client *client.QQClient, event *message.GroupMessage) *groupMessageEventPerformer {
	s := &groupMessageEventPerformer{
		Client: client,
		Event:  event,
	}
	s.strategyHandler = []func() bool{
		s.strategyMsg,
	}
	return s
}

func init() {
	service.RegisterGroupMessageEvent(new_())
}

func (s *sGroupMessageEvent) Event(client *client.QQClient, event *message.GroupMessage) {
	newPerformer(client, event).DoEvent()
}

func (s *groupMessageEventPerformer) DoEvent() {
	var done bool

	// 循环顺序判断策略，如果已经处理则退出
	for _, do := range s.strategyHandler {
		done = do()
		if done {
			return
		}
	}
}

func (s *groupMessageEventPerformer) strategyMsg() bool {
	//msg := s.Event.ToString()
	//
	//switch {
	//case
	//}
	return false
}
