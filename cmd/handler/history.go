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

type HistoryHandler struct {
	keys []string
	bot  *tgbotapi.BotAPI
}

func NewHistoryHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI) Handler {
	return &HistoryHandler{
		keys: []string{"История ставок 📜"},
		bot:  bot,
	}
}

func (h *HistoryHandler) Keys() []string {
	return h.keys
}

func (h *HistoryHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "История ставок \n\nДинамо (Киев) - Шахтер (Донецк) 100$ ❌\nУсик - Рокки 50$ ✅")
	// msg.ReplyMarkup = categoriesMenuKeyboard
	return msg
}
