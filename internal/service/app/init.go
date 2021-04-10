package app

import (
	"go-gemini-server/internal/app/geminiServer"
	"go-gemini-server/internal/service/log"
)

func Init() {
	geminiServer.Init()
	log.Info("appService init successfully")
}
