package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/fireferretsbet/tg-bot/internal/serverenv"
)

var topUpMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1 USDT"),
		tgbotapi.NewKeyboardButton("2 USDT"),
		tgbotapi.NewKeyboardButton("5 USDT"),
		tgbotapi.NewKeyboardButton("10 USDT"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("20 USDT"),
		tgbotapi.NewKeyboardButton("50 USDT"),
		tgbotapi.NewKeyboardButton("100 USDT"),
		tgbotapi.NewKeyboardButton("Назад ⬅️"),
	),
)

type TopUpHandler struct {
	GenericHandler
}

func NewTopUpHandler(env *serverenv.ServerEnv) Handler {
	return &TopUpHandler{
		GenericHandler{
			keys: []string{"Пополнить 💳"},
			env:  env,
		},
	}
}

func (h *TopUpHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите сумму пополнения")
	msg.ReplyMarkup = topUpMenuKeyboard
	return msg
}

func (h *TopUpHandler) GetDialogContext() string {
	return "top_up_success"
}

func (h *TopUpHandler) GetPreviousRoute() string {
	return "balance"
}
