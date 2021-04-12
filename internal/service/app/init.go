package app

import (
	"github.com/lmx-Hexagram/gemini-server/internal/app/geminiServer"
	"github.com/lmx-Hexagram/gemini-server/internal/service/log"
)

func Init() {
	geminiServer.Init()
	log.Info("appService init successfully")
}
