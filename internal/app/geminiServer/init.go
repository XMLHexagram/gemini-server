package geminiServer

import (
	"github.com/lmx-Hexagram/gemini-server/internal/service/config"
	"github.com/lmx-Hexagram/gemini-server/pkg/gemini"
)

var e *gemini.Engine

func Init() {
	tls := config.ProvideTLS()
	geminiConfig := config.ProvideGemini()
	engine, err := gemini.New(tls.CertFile, tls.KeyFile, geminiConfig.DefaultLang)
	if err != nil {
		panic(err)
	}
	engine.AutoRedirect = geminiConfig.AutoRedirect
	engine.AutoRedirectUrl = geminiConfig.AutoRedirectUrl
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
