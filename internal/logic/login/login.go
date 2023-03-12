package login

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/template-single/internal/dao"
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
	//bot := client.NewClient(s.number, s.psw)
	dao.AA()

	return nil
}
