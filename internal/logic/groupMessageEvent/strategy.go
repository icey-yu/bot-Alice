package groupMessageEvent

import (
	"bot-Alice/internal/service"
	"bot-Alice/internal/utils"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
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
		s.strategyMsg, // 消息 策略
	}
	return s
}

func init() {
	service.RegisterGroupMessageEvent(new_())
}

func (s *sGroupMessageEvent) Event(client *client.QQClient, event *message.GroupMessage) {
	ctx := gctx.New()
	g.Log().Infof(ctx, "收到群聊消息：GroupCode：%d,SenderUin：%d，msg：%s", event.GroupCode, event.Sender.Uin, event.ToString())
	err := newPerformer(client, event).DoEvent()
	if err != nil {
		// 该干嘛捏
		g.Log().Errorf(ctx, "群聊消息处理失败：%+v", err)
	}
	g.Log().Infof(ctx, "处理完毕")
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

	switch {
	case utils.IsAtRobotGroupStr(s.Client, s.Event): // @事件
		return s.atEvent(msg)
	}
	return false, nil
}
