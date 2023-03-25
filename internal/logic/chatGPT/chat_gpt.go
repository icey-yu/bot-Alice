package chatGPT

import (
	"bot-Alice/internal/consts"
	"bot-Alice/internal/global"
	"bot-Alice/internal/model"
	"bot-Alice/internal/service"
	"bot-Alice/internal/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Mrs4s/MiraiGo/message"
	"io"
	"sync"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/redis/go-redis/v9"
)

const (
	chatPath = "/api/ask"

	overTime              = time.Minute * 2               // 超时两分钟
	groupRedisKeyFormat   = "botAlice:chatGPT:group:%d"   // 群聊的redis会话消息保存key
	privateRedisKeyFormat = "botAlice:chatGPT:private:%d" // 私聊的redis会话消息保存key

	tryTimes    = 10              // 最多尝试次数。半分钟刷新时间，是很正常的
	refreshTime = time.Second * 6 // 重试的等待时间
)

var (
	// mutex
	mutex = sync.Mutex{}
	ch    = make(chan bool, 1000) // 用于计数
)

type (
	sChatGPT struct {
		address string
		key     string
	}
)

func new_() *sChatGPT {
	ctx := context.Background()
	return &sChatGPT{
		address: "http://" + g.Cfg().MustGet(ctx, "chatGPT.address").String(),
		key:     g.Cfg().MustGet(ctx, "chatGPT.key").String(),
	}
}

func init() {
	service.RegisterChatGPT(new_())
}

func (s *sChatGPT) GroupChat(groupMessage *message.GroupMessage, msg string) (string, error) {
	code := groupMessage.GroupCode
	replyMsg := message.NewReply(groupMessage)
	// 锁
	ch <- true
	if ok := mutex.TryLock(); !ok {
		// 有人正在使用
		sendMsg := utils.BuildTextMessage(fmt.Sprintf("有%d条消息在您之前使用chatGPT，稍后将为您重新调用~", len(ch)-1)).Append(replyMsg)
		global.Alice.SendGroupMessage(code, sendMsg)
		mutex.Lock()
	} else {
		// 发送消息表示已经接收到了请求
		sendMsg := message.NewSendingMessage()
		textMsg := message.NewText("让我想想呢🤔...") // 单独🤔 不支持reply消息，哈软！
		sendMsg.Elements = append(sendMsg.Elements, replyMsg, textMsg)
		global.Alice.SendGroupMessage(code, sendMsg)
	}
	defer func() {
		mutex.Unlock() // 解锁
		<-ch
	}()

	return s.chat(consts.Group, code, msg)
}

func (s *sChatGPT) PrivateChat(code int64, msg string) (string, error) {
	return s.chat(consts.Private, code, msg)
}

func (s *sChatGPT) chat(type_ int, code int64, msg string) (string, error) {
	msgData, err := s.getMsgData(type_, code)
	if err != nil {
		return "", gerror.Wrapf(err, "获取会话消息失败")
	}
	msgData.Content = msg

	// 添加重试
	reTryCount := 0
	var resp *model.ChatGPTResp
	for reTryCount < tryTimes {
		resp, err = s.sendChat(msgData)
		if err == nil {
			break
		}
		// 说明有错
		g.Log().Errorf(gctx.New(), err.Error())
		reTryCount++
		time.Sleep(refreshTime) // 等待页面刷新
	}
	if err != nil {
		return "", gerror.Wrapf(err, "发送聊天请求失败")
	}

	msgData.ConversationId = resp.ConversationId
	msgData.ParentId = resp.ResponseId
	err = s.saveMsgData(type_, code, *msgData)
	if err != nil {
		return "", gerror.Wrapf(err, "存储会话消息失败")
	}
	return resp.Content, nil
}

// sendChat 发送chat
func (s *sChatGPT) sendChat(msgData *model.ChatGPTMessage) (*model.ChatGPTResp, error) {
	jsData, err := json.Marshal(msgData)
	if err != nil {
		return nil, gerror.Wrapf(err, "序列化msgData失败：%s", msgData)
	}
	readerBody := bytes.NewReader(jsData)

	url := s.address + chatPath
	header := map[string]string{
		consts.Authorization: s.key,
	}

	response, err := utils.Post(url, readerBody, header, overTime)
	if err != nil {
		return nil, gerror.Wrapf(err, "chat请求失败")
	}
	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, gerror.Wrapf(err, "读取body失败")
	}

	var chatResp model.ChatGPTResp
	err = json.Unmarshal(respBody, &chatResp)
	if err != nil {
		return nil, gerror.Wrapf(err, "反序列化失败:", string(respBody))
	}

	// 如果chat返回错误
	if chatResp.Error != "" {
		return nil, gerror.Wrapf(gerror.New(chatResp.Error), "chat失败：返回体：%s", string(respBody))
	}

	return &chatResp, nil
}

// saveMsgData 存储聊天会话信息
func (s *sChatGPT) saveMsgData(type_ int, uin int64, msg model.ChatGPTMessage) error {
	ctx := gctx.New()
	var key string
	switch type_ {
	case consts.Group:
		key = fmt.Sprintf(groupRedisKeyFormat, uin)
	case consts.Private:
		key = fmt.Sprintf(privateRedisKeyFormat, uin)
	default:
		return gerror.New("不支持的类型")
	}

	msg.Content = "" // 减少不必要的存储

	jsData, err := json.Marshal(msg)
	if err != nil {
		return gerror.Wrapf(err, "序列化失败:%+v", msg)
	}
	err = global.Redis.Set(ctx, key, jsData, 0).Err()
	if err != nil {
		return gerror.Wrapf(err, "存储数据失败：", jsData)
	}
	return nil
}

// getMsgData 获取redis中保存的聊天会话信息
func (s *sChatGPT) getMsgData(type_ int, uin int64) (*model.ChatGPTMessage, error) {
	ctx := gctx.New()
	var key string
	switch type_ {
	case consts.Group:
		key = fmt.Sprintf(groupRedisKeyFormat, uin)
	case consts.Private:
		key = fmt.Sprintf(privateRedisKeyFormat, uin)
	default:
		return nil, gerror.New("不支持的类型")
	}
	result, err := global.Redis.Get(ctx, key).Result()
	switch err {
	case nil:
	// 忽略
	case redis.Nil:
		// 没有数据，新建一个默认数据返回
		return &model.ChatGPTMessage{}, nil
	default:
		return nil, gerror.Wrapf(err, "redis获取数据失败：key：", key)
	}
	var res model.ChatGPTMessage
	err = json.Unmarshal([]byte(result), &res)
	if err != nil {
		return nil, gerror.Wrapf(err, "反序列化失败：", result)
	}
	return &res, nil
}
