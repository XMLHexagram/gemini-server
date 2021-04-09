package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"path"
)

func (s *service) level() zap.AtomicLevel {
	var level zapcore.Level
	switch s.Log.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.DebugLevel
	}
	return zap.NewAtomicLevelAt(level)
}

func (s *service) customEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func (s *service) logRotateOptions() zap.Option {
	return zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   s.logRotateFileName(),
			MaxSize:    s.Log.Rotate.MaxSize,
			MaxAge:     s.Log.Rotate.MaxAge,
			MaxBackups: s.Log.Rotate.MaxBackups,
			LocalTime:  false,
			Compress:   false,
		})
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(s.customEncoderConfig()),
			w,
			s.level(),
		)
		return zapcore.NewTee(c, core)
	})
}

func (s *service) logRotateFileName() (fileName string) {
	for _, v := range s.Rotate.Filename {
		fileName = path.Join(fileName, v)
	}
	return fileName
}
