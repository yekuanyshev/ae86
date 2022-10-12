package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type ExtractContextFunc func(ctx context.Context) (key string, value interface{})

var (
	Log *zap.Logger
)

func LogCtx(ctx context.Context, f ExtractContextFunc) *zap.Logger {
	return Log.With(zap.Any(f(ctx)))
}

func Init() {
	encoder := getEncoder()
	writeSyncer := zapcore.AddSync(os.Stdout)
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	Log = zap.New(core)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
