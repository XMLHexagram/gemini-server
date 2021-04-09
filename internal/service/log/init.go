package log

import (
	"go-gemini-server/internal/service/config"
	"go.uber.org/zap"
)

var logService *service

type service struct {
	*zap.Logger        `wire:"-"`
	*zap.SugaredLogger `wire:"-"`
	config.Log
}

func Init() {
	s := initDep()

	var cfg = zap.Config{
		Level:             s.level(),
		Development:       s.Log.Development,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     s.customEncoderConfig(),
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
		InitialFields:     nil,
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	if s.Log.LogRotate == true {
		logger = logger.WithOptions(s.logRotateOptions())
	}

	defer logger.Sync()

	s.SugaredLogger = logger.Sugar()
	s.Logger = logger

	logService = s

	logService.Logger.Info("logService init successfully")
}
