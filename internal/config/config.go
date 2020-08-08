package config

import (
	"fmt"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

type Config struct {
	ApiToken string `envconfig:"API_TOKEN"`
	LogLevel string
	Port     int
	Debug    bool
}

func NewConfig() *Config {
	var config Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Can not read current directory %v", err)
	}
	viper.AddConfigPath(fmt.Sprintf("%s/../internal/config", path))
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	config.LogLevel = viper.GetString("logLevel")
	config.Port = viper.GetInt("port")
	config.Debug = viper.GetBool("debug")

	err = envconfig.Process("", &config)
	if err != nil {
		log.Fatalf("error loading environment variables: %v", err)
	}

	return &config
}
