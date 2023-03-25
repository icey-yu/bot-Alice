package utils

import (
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/gogf/gf/v2/errors/gerror"
)

// GetBotGroupDisplayName 获取机器人群昵称。如果没有昵称（CardName）则返回NickName
func GetBotGroupDisplayName(bot *client.QQClient, groupCode int64) (string, error) {
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

// RemoveGroupAtStr 去除string类型的群聊消息中的@
func RemoveGroupAtStr(bot *client.QQClient, groupCode int64, msg string) (string, error) {
	cardName, err := GetBotGroupDisplayName(bot, groupCode)
	if err != nil {
		return "", gerror.Wrap(err, "获取群昵称失败")
	}
	return RemoveAtStr(cardName, msg), nil
}
