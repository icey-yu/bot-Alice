package utils

import "github.com/Mrs4s/MiraiGo/message"

// RemoveAtEle 去除elements类型的消息中的@
func RemoveAtEle(code int64, messages []message.IMessageElement) []message.IMessageElement {
	for index, msg := range messages {
		if msg.Type() == message.At {
			atMsg := msg.(*message.AtElement)
			if atMsg.Target == code {
				SliceDelete(messages, index)
			}
		}
	}
	return messages
}

// RemoveReplyAt 在回复消息中移除多余的at
func RemoveReplyAt(messages []message.IMessageElement) []message.IMessageElement {
	for index, msg := range messages {
		if msg.Type() == message.Reply {
			if index < 2 { // 由qq群聊发送的reply消息，reply是第三个元素。
				return messages
			}
			if messages[index-1].Type() == message.Text && messages[index-2].Type() == message.At {
				// 当reply的前两个元素分别人At和内容为" "的text元素时，进行移除
				textMsg := messages[index-1].(*message.TextElement)
				if textMsg.Content == " " {
					if index == 2 {
						return append(messages[2:])
					}
					return append(messages[:index-3], messages[index:]...)
				}
			}
		}
	}
	return messages
}
