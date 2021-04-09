package config

import (
	"github.com/spf13/viper"
	"path"
)

var configService *service

type service struct {
	Config
}

func Init() {
	s := new(service)
	viper.AddConfigPath(path.Join("configs"))
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&s.Config)
	if err != nil {
		panic(err)
	}

	configService = s
}
