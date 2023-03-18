package consts

import "github.com/gogf/gf/v2/errors/gerror"

var (
	ErrChatIsLocked = gerror.New("有会话正在调用ChatGPT")
)

