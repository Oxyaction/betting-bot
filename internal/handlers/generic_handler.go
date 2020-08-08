package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/fireferretsbet/tg-bot/internal/user"
)

type GenericHandler struct {
	keys       []string
	bot        *tgbotapi.BotAPI
	userStates map[int]*user.UserState
}

func (h *GenericHandler) Keys() []string {
	return h.keys
}

func (h *GenericHandler) GetDialogContext() string {
	return ""
}

func (h *GenericHandler) GetPreviousRoute() string {
	return ""
}
