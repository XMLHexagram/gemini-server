package service

import (
	"go-gemini-server/internal/app/geminiServer"
	"go-gemini-server/internal/service/app"
	"go-gemini-server/internal/service/config"
	"go-gemini-server/internal/service/log"
)

func Init() {
	config.Init()
	log.Init()
	app.Init()
	geminiServer.Init()
}

func Run() {
	geminiServer.Run()
}
