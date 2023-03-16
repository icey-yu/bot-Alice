package utils

import (
	"github.com/Mrs4s/MiraiGo/message"
)

// BuildSendingMessage 构建发送信息
func BuildSendingMessage(msg string) *message.SendingMessage {
	sendingMessage := message.NewSendingMessage()
	sendingMessage.Elements = append(sendingMessage.Elements, message.NewText(msg))
	return sendingMessage
}
