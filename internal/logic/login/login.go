package login

import (
	"bot-Alice/internal/global"
	"bot-Alice/internal/service"
	"context"
	"fmt"
	"os"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
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

func new_() *sLogin {
	ctx := context.Background()
	return &sLogin{
		number:   g.Cfg().MustGet(ctx, "qqAccount.number").Int64(),
		psw:      g.Cfg().MustGet(ctx, "qqAccount.psw").String(),
		filePath: g.Cfg().MustGet(ctx, "filePath").String(),
	}
}

func init() {
	service.RegisterLogin(new_())
}

// Login 登录
func (s *sLogin) Login() error {
	ctx := context.Background()

	if global.Alice != nil && global.Alice.Online.Load() {
		g.Log().Info(ctx, "已经登录，无需重新登录")
		return nil
	}

	g.Log().Printf(ctx, "开始登录")
	global.Alice = client.NewClient(s.number, s.psw)
	global.Alice.UseDevice(client.GenRandomDevice())
	global.Alice.AllowSlider = true

	err := s.tokenLogin(ctx)
	if err != nil {
		return gerror.Wrapf(err, "登录失败")
	}
	g.Log().Printf(ctx, "登录完毕")
	return nil
}

// tokenLogin 默认token登录
func (s *sLogin) tokenLogin(ctx context.Context) error {
	// 读取token
	token, err := os.ReadFile(s.filePath + tokenTXTName)
	if err != nil {
		g.Log().Error(ctx, gerror.Wrapf(err, "读文件失败"))
	}
	if len(token) != 0 {
		err = global.Alice.TokenLogin(token)
		if err != nil {
			g.Log().Error(ctx, gerror.Wrapf(err, "token登录失败"))
		} else {
			return nil
		}
	}
	// 登录未成功，普通登录
	err = s.commonLogin(ctx)
	if err != nil {
		return gerror.Wrapf(err, "账号密码登录失败")
	}

	return nil
}

// commonLogin 账号密码登录
func (s *sLogin) commonLogin(ctx context.Context) error {
	login, err := global.Alice.Login()
	if err != nil {
		return gerror.Wrapf(err, "登录失败")
	}
	if login.Error.String() != "" {
		return gerror.New(login.Error.String())
	}
	g.Log().Printf(ctx, "登录状态:%t", login.Success)
	// 写token
	err = os.WriteFile(fmt.Sprintf("%stoken.txt", s.filePath), global.Alice.GenToken(), 0777)
	if err != nil {
		g.Log().Error(ctx, "写token失败。")
	}

	return nil
}
