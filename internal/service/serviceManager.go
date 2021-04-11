package service

import (
	"go-gemini-server/internal/service/app"
	"go-gemini-server/internal/service/config"
	"go-gemini-server/internal/service/log"
)

func Init() {
	config.Init()
	log.Init()
	app.Init()
}

func Run() {
	app.Run()
}
