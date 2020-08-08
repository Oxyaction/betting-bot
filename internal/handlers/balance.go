package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
	"gitlab.com/fireferretsbet/tg-bot/internal/user"
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

func NewBalanceHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI, userStates map[int]*user.UserState) Handler {
	return &BalanceHandler{
		GenericHandler{
			keys:       []string{"Баланс 🏦", "top_up_success"},
			bot:        bot,
			userStates: userStates,
		},
	}
}

func (h *BalanceHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	var text string
	// top_up_success
	if h.userStates[update.Message.From.ID].ContextRoute == "top_up_success" {
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
