package utils

import (
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/gogf/gf/v2/errors/gerror"
)

// GetGroupCardName 获取机器人群聊中的名称
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
	return self.CardName, nil
}
