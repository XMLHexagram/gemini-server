package log

import "go.uber.org/zap"

func Info(msg string, fields ...zap.Field) {
	logService.WithOptions(zap.AddCallerSkip(1)).Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logService.WithOptions(zap.AddCallerSkip(1)).Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logService.WithOptions(zap.AddCallerSkip(1)).Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logService.WithOptions(zap.AddCallerSkip(1)).Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	logService.WithOptions(zap.AddCallerSkip(1)).DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logService.WithOptions(zap.AddCallerSkip(1)).Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logService.WithOptions(zap.AddCallerSkip(1)).Fatal(msg, fields...)
}

func ProvideSugaredLogger() *zap.SugaredLogger {
	return logService.SugaredLogger
}

func ProvideLogger() *zap.Logger {
	return logService.Logger
}
