package login

import (
	"context"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkg/errors"
	"os"
)

type (
	sLogin struct {
		number int64
		psw    string
	}
)

func _new() *sLogin {
	ctx := context.Background()
	return &sLogin{
		number: g.Cfg().MustGet(ctx, "qqAccount.number").Int64(),
		psw:    g.Cfg().MustGet(ctx, "qqAccount.psw").String(),
	}
}

func (s *sLogin) Login() error {
	bot := client.NewClient(s.number, s.psw)
	bot.UseDevice(client.GenRandomDevice())
	bot.AllowSlider = true

	//res, err := bot.Login()
	//println(bot.GenToken())
	//os.WriteFile("token.txt", bot.GenToken(), 0777)

	token, err := os.ReadFile("token.txt")
	if err != nil {
		return errors.Wrapf(err, "读文件失败")
	}
	err = os.WriteFile("device.txt", bot.Device().ToJson(), 0777)
	if err != nil {
		return errors.Wrapf(err, "写文件失败")
	}

	err = bot.TokenLogin(token)
	if err != nil {
		return
	}

	return nil
}
