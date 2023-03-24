package groupMessageEvent

import (
	"bot-Alice/internal/consts"
	"bot-Alice/internal/service"
	"bot-Alice/internal/utils"
	"github.com/Mrs4s/MiraiGo/message"

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
		msgExChat := utils.RemoveChat(msgExAt)
		err := s.callChatGPT(msgExChat)
		if err != nil {
			return true, gerror.Wrapf(err, "调用ChatGPT失败")
		}
	}
	return false, nil
}

// callChatGPT 调用chatGPT
func (s *groupMessageEventPerformer) callChatGPT(msg string) error {
	// 开始调用chatGPT
	res, err := service.ChatGPT().GroupChat(s.Event, msg)
	if err != nil {
		if gerror.Is(err, consts.ErrChatIsLocked) {
			// 因为有人正在聊天而失败
			sendMsg := utils.BuildTextMessage("有其他人正在聊天呢，请稍等🙃")
			s.Client.SendGroupMessage(s.Event.GroupCode, sendMsg)
		} else {
			sendMsg := utils.BuildTextMessage("和chatGPT聊天失败了呢😔")
			s.Client.SendGroupMessage(s.Event.GroupCode, sendMsg)
		}
		return gerror.Wrapf(err, "调用chatGPT失败")
	}
	sendingMessage := message.NewSendingMessage()
	reply := message.NewReply(s.Event)
	at := message.NewAt(s.Event.Sender.Uin) // at并没有起效果qaq
	msgStr := message.NewText(res)
	sendingMessage.Elements = append(sendingMessage.Elements, at, reply, at, msgStr)
	s.Client.SendGroupMessage(s.Event.GroupCode, sendingMessage)
	return nil
}

// praiseEvent 夸赞
func (s *groupMessageEventPerformer) praiseEvent() {
	sendMsg := utils.BuildTextMessage("高性能ですから!")
	s.Client.SendGroupMessage(s.Event.GroupCode, sendMsg)
}
