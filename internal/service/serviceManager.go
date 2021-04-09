package service

import (
	"go-gemini-server/internal/service/app"
	"go-gemini-server/internal/service/config"
	"go-gemini-server/internal/service/db"
	"go-gemini-server/internal/service/httpEngine"
	"go-gemini-server/internal/service/log"
	"go-gemini-server/internal/service/toolkit"
)

func Init() {
	config.Init()
	log.Init()
	toolkit.Init()
	db.Init()
	httpEngine.Init()
	app.Init()
}

func Run() {
	httpEngine.Run()
}
