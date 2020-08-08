package main

import (
	"context"
	"net/http"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
	"gitlab.com/fireferretsbet/tg-bot/internal/handlers"
	"gitlab.com/fireferretsbet/tg-bot/internal/logger"
)

func main() {
	config := config.NewConfig()
	log := logger.NewLogger(config)

	log.WithFields(logrus.Fields{
		"port": config.Port,
	}).Info("server started")

	bot, err := tgbotapi.NewBotAPI(config.ApiToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe("0.0.0.0:"+strconv.Itoa(config.Port), nil)

	h := handlers.NewUpdateHandler(log, config, bot)
	ctx := context.Background()
	for update := range updates {
		go h.Handle(update, ctx)
	}
}
