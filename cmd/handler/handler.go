package handler

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
)

type UpdateHandler struct {
	log    *logrus.Logger
	config *config.Config
	bot    *tgbotapi.BotAPI
}

func NewHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI) *UpdateHandler {
	return &UpdateHandler{
		log,
		config,
		bot,
	}
}

func (h *UpdateHandler) Handle(update tgbotapi.Update, ctx context.Context) {
	if update.Message == nil {
		return
	}

	h.log.WithFields(logrus.Fields{
		"user_id":  update.Message.From.ID,
		"username": update.Message.From.UserName,
		"text":     update.Message.Text,
	}).Info("message accepted")

	if update.Message.IsCommand() {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		switch update.Message.Command() {
		case "start":
			msg.Text = "Welcome to FireFerrets betting bot. 🔥🐹🤑"
		default:
			msg.Text = "I don't know that command"
		}
		h.bot.Send(msg)
	} else {

	}
}
