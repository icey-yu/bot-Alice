package utils

import (
	"bot-Alice/internal/consts"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/gogf/gf/v2/text/gstr"
)

// RemoveAt 去除@
func RemoveAt(name, msg string) string {
	msg = gstr.Replace(msg, "@"+name+" ", "", -1)
	msg = gstr.Replace(msg, "@"+name, "", -1)
	return msg
}

// RemoveChat 去除 chat:
func RemoveChat(msg string) string {
	removeList := consts.ChatGPTChat
	for _, re := range removeList {
		msg = gstr.TrimLeft(msg, re)
	}
	return msg
}

// BuildTextMessage 构建发送文本信息
func BuildTextMessage(msg string) *message.SendingMessage {
	sendingMessage := message.NewSendingMessage()
	SendMsgAppends(sendingMessage, message.NewText(msg))
	return sendingMessage
}

// SendMsgAppends 添加需要发送的消息
func SendMsgAppends(sendMsg *message.SendingMessage, messages ...message.IMessageElement) {
	// 回复消息中，如果有at，需要at两次
	sendMsg.Elements = append(sendMsg.Elements, messages...)
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
