package zaplog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitConsoleLog() *zap.Logger {
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(zapcore.EncoderConfig{}), os.Stdout, zapcore.InfoLevel)
	return zap.New(core)
}
