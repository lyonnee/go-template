package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/lyonnee/go-template/pkg/log"
	"go.uber.org/zap"
)

func Recovery(ctx context.Context, reqCtx *app.RequestContext, err interface{}, stack []byte) {
	log.Panic("[Recovery from panic]", zap.Any("err", err), zap.ByteString("stack", stack))
	reqCtx.AbortWithStatus(consts.StatusInternalServerError)
}
