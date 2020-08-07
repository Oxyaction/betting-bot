package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	ApiToken string `envconfig:"API_TOKEN"`
	logLevel string
	port     int
	debug    bool
	foo      int `envconfig:"FOO"`
}

var config Config

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Can not read current directory %v", err)
	}
	viper.AddConfigPath(path)
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	config.logLevel = viper.GetString("logLevel")
	config.port = viper.GetInt("port")
	config.debug = viper.GetBool("debug")

	err = envconfig.Process("", &config)
	if err != nil {
		log.Fatalf("error loading environment variables: %v", err)
	}

	fmt.Printf("%#v\n", config)
}

// Password: viper.GetString("db.password"),

func initLogger() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	level, err := log.ParseLevel(config.logLevel)
	if err != nil {
		level = log.InfoLevel
	}

	log.SetLevel(level)
}

func main() {
	initConfig()
	initLogger()

	log.WithFields(log.Fields{
		"port": config.port,
	}).Info("server started")

	bot, err := tgbotapi.NewBotAPI(config.ApiToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = config.debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// _, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://www.google.com:8443/"+config.ApiToken, "cert.pem"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe("0.0.0.0:"+strconv.Itoa(config.port), nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}
