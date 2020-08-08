package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type GenericHandler struct {
	keys []string
	bot  *tgbotapi.BotAPI
}

func (h *GenericHandler) Keys() []string {
	return h.keys
}

func (h *GenericHandler) GetDialogContext() string {
	return ""
}
