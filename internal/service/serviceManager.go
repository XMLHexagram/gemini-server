package service

import (
	"github.com/lmx-Hexagram/gemini-server/internal/service/app"
	"github.com/lmx-Hexagram/gemini-server/internal/service/config"
	"github.com/lmx-Hexagram/gemini-server/internal/service/log"
)

func Init() {
	config.Init()
	log.Init()
	app.Init()
}

func Run() {
	app.Run()
}
