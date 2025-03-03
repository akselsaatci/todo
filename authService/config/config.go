package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	DATABASE_HOST     string `mapstructure:"DATABASE_HOST"`
	DATABASE_PORT     string `mapstructure:"DATABASE_PORT"`
	DATABASE_USER     string `mapstructure:"DATABASE_USER"`
	DATABASE_PASSWORD string `mapstructure:"DATABASE_PASSWORD"`
	DATABASE_NAME     string `mapstructure:"DATABASE_NAME"`
	SERVER_PORT       string `mapstructure:"SERVER_PORT"`
	JWT_SECRET_KEY    string `mapstructure:"JWT_SECRET_KEY"`
}

func ReadConfig() *AppConfig {
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config") // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}

	var appConfig AppConfig

	err = viper.Unmarshal(&appConfig)

	if err != nil {
		panic(fmt.Errorf("fatal error unmarshalling config file: %w", err))
	}

	return &appConfig

}
