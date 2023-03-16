package groupMessageEvent

import (
	"bot-Alice/internal/service"
	"bot-Alice/internal/utils"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	sGroupMessageEvent struct{}

	// groupMessageEventPerformer 具体的事件执行
	groupMessageEventPerformer struct {
		Client          *client.QQClient
		Event           *message.GroupMessage
		strategyHandler []func() (bool, error)
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
	s.strategyHandler = []func() (bool, error){
		s.strategyMsg,
	}
	return s
}

func init() {
	service.RegisterGroupMessageEvent(new_())
}

func (s *sGroupMessageEvent) Event(client *client.QQClient, event *message.GroupMessage) {
	err := newPerformer(client, event).DoEvent()
	if err != nil {
		// 该干嘛捏
	}
}

func (s *groupMessageEventPerformer) DoEvent() error {

	// 循环顺序判断策略，如果已经处理则退出
	for _, do := range s.strategyHandler {
		done, err := do()
		if done || err != nil {
			return err
		}
	}
	return nil
}

func (s *groupMessageEventPerformer) strategyMsg() (bool, error) {
	msg := s.Event.ToString()

	isAt, err := utils.IsAtRobotGroup(s.Client, s.Event.GroupCode, msg)
	if err != nil {
		return false, gerror.Wrap(err, "获取是否@bot失败")
	}

	switch {
	case isAt: // @事件

	}
	return false, nil
}
