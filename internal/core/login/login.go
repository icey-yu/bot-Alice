package login

import (
	"bot-Alice/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

// Init 登录初始化
func Init() error {
	err := service.Login().Login()
	if err != nil {
		return gerror.Wrap(err, "初始登录失败")
	}
	return nil
}
