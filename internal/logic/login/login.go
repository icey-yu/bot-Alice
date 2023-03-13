package login

import (
	"bot-Alice/internal/global"
	"context"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"os"
)

const (
	tokenTXTName = "token.txt"
)

type (
	sLogin struct {
		number   int64
		psw      string
		filePath string
	}
)

func _new() *sLogin {
	ctx := context.Background()
	return &sLogin{
		number:   g.Cfg().MustGet(ctx, "qqAccount.number").Int64(),
		psw:      g.Cfg().MustGet(ctx, "qqAccount.psw").String(),
		filePath: g.Cfg().MustGet(ctx, "filePath").String(),
	}
}

func (s *sLogin) Login() error {
	ctx := context.Background()
	global.Alice = client.NewClient(s.number, s.psw)

	global.Alice.UseDevice(client.GenRandomDevice())
	global.Alice.AllowSlider = true

	err := s.tokenLogin()
	if err != nil {
		g.Log().Printf(ctx, err.Error())
	}
	res, err := bot.Login()
	println(bot.GenToken())
	os.WriteFile("token.txt", bot.GenToken(), 0777)

	err = os.WriteFile("device.txt", bot.Device().ToJson(), 0777)
	if err != nil {
		return gerror.Wrapf(err, "写文件失败")
	}

	return nil
}

func (s *sLogin) tokenLogin() error {
	token, err := os.ReadFile(s.filePath + tokenTXTName)
	if err != nil {
		return gerror.Wrapf(err, "读文件失败")
	}
	err = global.Alice.TokenLogin(token)
	if err != nil {
		return gerror.Wrapf(err, "token登录失败")
	}
	return nil
}
