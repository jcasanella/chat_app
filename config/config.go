package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init viper library with location of the config file
func Init(env string) {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath(".")
	config.AddConfigPath("config/")

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("error on parsing configuration file. %v", err)
	}
}

// GetConfig returns a Viper object
func GetConfig() *viper.Viper {
	return config
}
