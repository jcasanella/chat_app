package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath(".")

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("error on parsing configuration file. %v", err)
	}
}

func GetConfig() *viper.Viper {
	return config
}
