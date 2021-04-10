package geminiServer

import (
	"go-gemini-server/internal/service/config"
	"go-gemini-server/pkg/gemini"
)

var e *gemini.Engine

func Init() {
	tls := config.ProvideTLS()
	engine, err := gemini.New(tls.CertFile, tls.KeyFile)
	if err != nil {
		panic(err)
	}

	geminiConfig := config.ProvideGemini()
	for _, v := range geminiConfig.Dir {
		engine.HandleDir(v.Router, v.Path, v.Index)
	}

	for _, v := range geminiConfig.File {
		engine.HandleFile(v.Router, v.Path)
	}

	for _, v := range geminiConfig.Proxy {
		engine.HandleProxy(v.Router, v.URL)
	}

	e = engine
}

func Run() error {
	server := config.ProvideServer("gemini")
	err := e.Run(server.Port)
	if err != nil {
		return err
	}
	return nil
}
