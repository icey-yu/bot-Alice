package utils

import (
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/gogf/gf/v2/errors/gerror"
)

// GetGroupCardName 获取群昵称
func GetGroupCardName(bot *client.QQClient, groupCode int64) (string, error) {
	info, err := bot.GetGroupInfo(groupCode)
	if err != nil {
		return "", gerror.Wrap(err, "获取群聊失败")
	}
	members, err := bot.GetGroupMembers(info)
	if err != nil {
		return "", gerror.Wrapf(err, "获取成员列表失败")
	}
	info.Members = members
	self := info.FindMember(bot.Uin)
	if self == nil {
		return "", gerror.New("没有找到自己QAQ")
	}

	if self.CardName == "" {
		self.CardName = bot.Nickname
	}
	return self.CardName, nil
}

// RemoveGroupAt 去除群聊消息中的@
func RemoveGroupAt(bot *client.QQClient, groupCode int64, msg string) (string, error) {
	cardName, err := GetGroupCardName(bot, groupCode)
	if err != nil {
		return "", gerror.Wrap(err, "获取群昵称失败")
	}
	return RemoveAt(cardName, msg), nil
}
