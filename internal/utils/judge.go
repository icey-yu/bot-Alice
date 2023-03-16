package utils

import (
	"fmt"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
)

// IsAtRobotGroup 群聊消息中是否@机器人
func IsAtRobotGroup(bot *client.QQClient, groupCode int64, msg string) (bool, error) {
	cardName, err := GetGroupCardName(bot, groupCode)
	if err != nil {
		return false, gerror.Wrapf(err, "获取botCardName失败")
	}
	return IsAt(cardName, msg), nil
}

// IsAt 是否@
func IsAt(name string, msg string) bool {
	atStr := fmt.Sprintf("@%s", name)
	return gstr.Contains(msg, atStr)
}

// IsPraise 是否夸赞
func IsPraise(msg string) {

}
