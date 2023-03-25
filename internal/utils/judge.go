package utils

import (
	"fmt"
	"github.com/Mrs4s/MiraiGo/message"
	"strings"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/dlclark/regexp2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
)

// IsAtRobotGroupStr 群聊消息中是否@机器人
func IsAtRobotGroupStr(bot *client.QQClient, groupCode int64, messages []message.IMessageElement, msg string) bool {
	ctx := gctx.New()
	cardName, err := GetBotGroupDisplayName(bot, groupCode)
	if err != nil {
		g.Log().Errorf(ctx, "获取botCardName失败:%v", err)
		return false
	}

	return IsAtStr(cardName, msg) || IsAtEle(bot.Uin, messages)
}

// IsAtStr 通过字符串判断是否@
func IsAtStr(name string, msg string) bool {
	atStr := fmt.Sprintf("@%s", name)
	return gstr.Contains(msg, atStr)
}

// IsAtEle 通过element判断是否@
func IsAtEle(code int64, messages []message.IMessageElement) bool {
	for _, msg := range messages {
		if msg.Type() == message.At {
			atMsg := msg.(*message.AtElement)
			if atMsg.Target == code {
				return true
			}
		}
	}
	return false
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

	msg = strings.TrimSpace(msg)
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
