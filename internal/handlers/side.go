package handlers

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/fireferretsbet/tg-bot/internal/serverenv"
)

var sideMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Back"),
		tgbotapi.NewKeyboardButton("Lay"),
		tgbotapi.NewKeyboardButton("Назад ⬅️"),
	),
)

type SideHandler struct {
	GenericHandler
}

func NewSideHandler(env *serverenv.ServerEnv) Handler {
	return &SideHandler{
		GenericHandler{
			keys: []string{"side"},
			env:  env,
		},
	}
}

func (h *SideHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	event := h.env.EventManager().EventByName(update.Message.Text)
	if event == nil {
		return tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Событие '%s' не найдено", update.Message.Text))
	}
	h.env.UserManager().SetEvent(update.Message.From.ID, event.ID)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите сторону")
	msg.ReplyMarkup = sideMenuKeyboard
	return msg
}

func (h *SideHandler) GetDialogContext() string {
	return "coeff"
}

func (h *SideHandler) GetPreviousRoute() string {
	return "categories"
}
