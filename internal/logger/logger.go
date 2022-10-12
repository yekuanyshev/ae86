package logger

import (
	"context"
	"github.com/supernova0730/ae86/internal/transport/utils"
	"github.com/supernova0730/ae86/pkg/logger"
	"go.uber.org/zap"
)

func LogWithCtx(ctx context.Context) *zap.Logger {
	return logger.LogCtx(ctx, extractMetaFromCtx)
}

func extractMetaFromCtx(ctx context.Context) (key string, value interface{}) {
	return "meta", utils.GetMeta(ctx)
}
