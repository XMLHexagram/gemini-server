package config

import (
	_ "embed"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
)

var configService *service

//go:embed config.toml.d
var DConfig string

type service struct {
	Config
}

func Init() {
	s := new(service)
	viper.AddConfigPath(path.Join("."))
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	fmt.Println(err)
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		f, err := os.Create("config.toml.d")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		_, err = f.Write([]byte(DConfig))
		if err != nil {
			panic(err)
		}
		fmt.Println("create config.toml.d, please rename to config.toml after configure")
		os.Exit(1)

	}
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&s.Config)
	if err != nil {
		panic(err)
	}

	configService = s
}
