package handler

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
)

// var balanceMenuKeyboard = tgbotapi.NewReplyKeyboard(
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("Спорт ⚽"),
// 		tgbotapi.NewKeyboardButton("Киберспорт 🎮"),
// 		tgbotapi.NewKeyboardButton("Политика 🏛️"),
// 		tgbotapi.NewKeyboardButton("Главное меню ⬅️"),
// 	),
// )

type BalanceHandler struct {
	keys []string
	bot  *tgbotapi.BotAPI
}

func NewBalanceHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI) Handler {
	return &BalanceHandler{
		keys: []string{"Баланс 🏦"},
		bot:  bot,
	}
}

func (h *BalanceHandler) Keys() []string {
	return h.keys
}

func (h *BalanceHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Баланс 🏦\n\nВыберите интересующий раздел.")
	// msg.ReplyMarkup = balanceMenuKeyboard
	return msg
}
