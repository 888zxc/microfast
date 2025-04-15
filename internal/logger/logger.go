package logger

import (
	"go.uber.org/zap"
)

var zapLogger *zap.Logger

func InitLogger() {
	if zapLogger == nil {
		logger, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}
		zapLogger = logger
	}
}

func L() *zap.Logger {
	if zapLogger == nil {
		InitLogger()
	}
	return zapLogger
}
