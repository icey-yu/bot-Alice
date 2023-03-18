package utils

import (
	"fmt"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/dlclark/regexp2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
)

// IsAtRobotGroup 群聊消息中是否@机器人
func IsAtRobotGroup(bot *client.QQClient, groupCode int64, msg string) bool {
	ctx := gctx.New()
	cardName, err := GetGroupCardName(bot, groupCode)
	if err != nil {
		g.Log().Errorf(ctx, "获取botCardName失败:%v", err)
		return false
	}
	return IsAt(cardName, msg)
}

// IsAt 是否@
func IsAt(name string, msg string) bool {
	atStr := fmt.Sprintf("@%s", name)
	return gstr.Contains(msg, atStr)
}

// IsPraise 是否夸赞
func IsPraise(msg string) bool {
	ctx := gctx.New()
	praiseFormat := []string{
		"^你?真?(可以|厉害|棒|不错)[！|!]?$",
		"^(?:very )?good[！|!]?$",
	}
	for _, format := range praiseFormat {
		match := regexp2.MustCompile(format, regexp2.IgnoreCase)
		isMatch, err := match.MatchString(msg)
		if err != nil {
			g.Log().Errorf(ctx, "正则匹配失败:%v", err)
			return false
		}
		if isMatch {
			return isMatch
		}
	}
	return false
}

// IsChatGPT 是否请求调用chatGPT。规则为:  chat:
func IsChatGPT(msg string) bool {
	// 开头是否是 chat:

	msg = gstr.ToLower(msg)
	ruleList := []string{"chat:", "chat："}
	for _, rule := range ruleList {
		if len(msg) < len(rule) {
			continue
		}
		if msg[:len(rule)] == rule {
			return true
		}
	}
	return false
}
