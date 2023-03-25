package utils

import (
	"bot-Alice/internal/consts"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/gogf/gf/v2/text/gstr"
)

// RemoveAtStr 根据字符串消息类型去除@
func RemoveAtStr(name, msg string) string {
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
