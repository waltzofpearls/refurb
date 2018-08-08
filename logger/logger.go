package logger

import "go.uber.org/zap"

type Logger interface {
	Sync() error
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
}

var global Logger

func Init() (err error) {
	global, err = zap.NewProduction()
	return
}

func Sync() error {
	return global.Sync()
}

func Info(msg string, fields ...zap.Field) {
	global.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	global.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	global.Fatal(msg, fields...)
}
