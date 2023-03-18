package utils

import (
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
	msg = gstr.TrimLeft(msg, "chat:")
	msg = gstr.TrimLeft(msg, "chat：")
	return msg
}


// BuildTextMessage 构建发送文本信息
func BuildTextMessage(msg string) *message.SendingMessage {
	sendingMessage := message.NewSendingMessage()
	sendingMessage.Elements = append(sendingMessage.Elements, message.NewText(msg))
	return sendingMessage
}
