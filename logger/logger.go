package logger

import (
	"go.uber.org/zap"
	"os"
)

var Logger *zap.Logger

func InitLogger() {
	logLevel := os.Getenv("LOG_LEVEL")

	var err error
	if logLevel == "debug" {
		Logger, err = zap.NewDevelopment()
	} else {
		Logger, err = zap.NewProduction()
	}
	if err != nil {
		panic("failed to initialize logger")
	}
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}
