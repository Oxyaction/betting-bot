package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Handler interface {
	Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig
	Keys() []string
	GetDialogContext() string
}
