package groupMessageEvent

import (
	"bot-Alice/internal/service"
	"bot-Alice/internal/utils"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (s *groupMessageEventPerformer) atEvent(msg string) (bool, error) {
	// 是否是夸赞~

	msgExAt, err := utils.RemoveGroupAt(s.Client, s.Event.GroupCode, msg)
	if err != nil {
		return false, gerror.Wrapf(err, "移除@失败")
	}
	switch {
	case utils.IsPraise(msgExAt):
		s.praiseEvent()
		return true, nil
	case utils.IsChatGPT(msgExAt):
		err := s.callChatGPT(msgExAt)
		if err != nil {
			return true, gerror.Wrapf(err, "调用ChatGPT失败")
		}
	}
	return false, nil
}

// callChatGPT 调用chatGPT
func (s *groupMessageEventPerformer) callChatGPT(msg string) error {
	sendMsg := utils.BuildTextMessage("🤔")
	s.Client.SendGroupMessage(s.Event.GroupCode, sendMsg)
	// 开始调用chatGPT
	res, err := service.ChatGPT().GroupChat(s.Event.GroupCode, msg)
	if err != nil {
		sendMsg := utils.BuildTextMessage("和chatGPT聊天失败了呢😔")
		s.Client.SendGroupMessage(s.Event.GroupCode, sendMsg)
		return gerror.Wrapf(err, "调用chatGPT失败")
	}
	sendMsg = utils.BuildTextMessage(res)
	s.Client.SendGroupMessage(s.Event.GroupCode, sendMsg)
	return nil
}

// praiseEvent 夸赞
func (s *groupMessageEventPerformer) praiseEvent() {
	sendMsg := utils.BuildTextMessage("高性能ですから!")
	s.Client.SendGroupMessage(s.Event.GroupCode, sendMsg)
}
