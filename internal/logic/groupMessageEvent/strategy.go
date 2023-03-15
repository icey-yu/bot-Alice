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
		Client *client.QQClient
		Event  *message.GroupMessage
	}
)

func new_() *sGroupMessageEvent {
	return &sGroupMessageEvent{}
}

func newPerformer(client *client.QQClient, event *message.GroupMessage) *groupMessageEventPerformer {
	return &groupMessageEventPerformer{
		Client: client,
		Event:  event,
	}
}

func init() {
	service.RegisterGroupMessageEvent(new_())
}

func (s *sGroupMessageEvent) Event(client *client.QQClient, event *message.GroupMessage) {
	newPerformer(client, event).DoEvent()
}

func (s *groupMessageEventPerformer) DoEvent() {
	var done bool
	msg := s.Event.ToString()
	groupCode := s.Event.GroupCode
	senderUin := s.Event.Sender.Uin

}

func (s *groupMessageEventPerformer) strategyMsg(msg string) bool {
	s.Client.Nickname
	switch msg {
	case
	}
}
