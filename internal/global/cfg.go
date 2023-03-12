package global

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	ctx      = context.Background()
	FilePath = g.Cfg().MustGet(ctx, "filePath").String()
)
