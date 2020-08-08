package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/fireferretsbet/tg-bot/internal/serverenv"
)

var sideMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Back"),
		tgbotapi.NewKeyboardButton("Lay"),
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
	h.env.UserManager().SetMatch(update.Message.From.ID, update.Message.Text)

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
