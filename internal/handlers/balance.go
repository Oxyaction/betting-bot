package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/fireferretsbet/tg-bot/internal/serverenv"
)

var balanceMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Пополнить 💳"),
		tgbotapi.NewKeyboardButton("Главное меню ⬅️"),
	),
)

type BalanceHandler struct {
	GenericHandler
}

func NewBalanceHandler(env *serverenv.ServerEnv) Handler {
	return &BalanceHandler{
		GenericHandler{
			keys: []string{"Баланс 🏦", "top_up_success"},
			env:  env,
		},
	}
}

func (h *BalanceHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	var text string
	// top_up_success
	if h.env.UserManager().GetContextRoute(update.Message.From.ID) == "top_up_success" {
		text = "Баланс успешно пополнен ✅\n\n"
	} else {
		text = "Баланс 🏦\n\n"
	}
	text += "Ваш текущий баланс: *100 $*."
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = balanceMenuKeyboard
	return msg
}
