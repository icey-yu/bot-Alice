package groupMessageEvent

import (
	"bot-Alice/internal/utils"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (s *groupMessageEventPerformer) atEvent(msg string) (bool, error) {
	// 是否是夸赞~
	praise, err := utils.IsPraise(msg)
	if err != nil {
		return false, gerror.Wrapf(err, "解析是否praise失败")
	}
	switch {
	case praise:
		s.praiseEvent()
		return true, nil
	}
	return false, nil
}

// praiseEvent 夸赞
func (s *groupMessageEventPerformer) praiseEvent() {
	sendMsg := utils.BuildSendingMessage("高性能ですから!")
	s.Client.SendGroupMessage(s.Event.GroupCode, sendMsg)
}
