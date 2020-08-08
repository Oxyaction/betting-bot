package handler

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
)

// var categoriesMenuKeyboard = tgbotapi.NewReplyKeyboard(
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("Спорт ⚽"),
// 		tgbotapi.NewKeyboardButton("Киберспорт 🎮"),
// 		tgbotapi.NewKeyboardButton("Политика 🏛️"),
// 		tgbotapi.NewKeyboardButton("Главное меню ⬅️"),
// 	),
// )

type AckedBetsHandler struct {
	keys []string
	bot  *tgbotapi.BotAPI
}

func NewAckedBetsHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI) Handler {
	return &AckedBetsHandler{
		keys: []string{"Подтвержденные ставки ✅"},
		bot:  bot,
	}
}

func (h *AckedBetsHandler) Keys() []string {
	return h.keys
}

func (h *AckedBetsHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Подтвержденные ставки ✅\n\nДинамо (Киев) - Шахтер (Донецк) 100$")
	// msg.ReplyMarkup = categoriesMenuKeyboard
	return msg
}
