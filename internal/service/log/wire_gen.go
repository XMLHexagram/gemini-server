// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package log

import (
	"github.com/lmx-Hexagram/gemini-server/internal/service/config"
)

// Injectors from wire.go:

func initDep() *service {
	log := config.ProvideLog()
	service2 := &service{
		Log: log,
	}
	return service2
}
