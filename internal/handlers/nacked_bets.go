package handlers

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

type NackedBetsHandler struct {
	GenericHandler
}

func NewNackedBetsHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI) Handler {
	return &NackedBetsHandler{
		GenericHandler{
			keys: []string{"Неподтвержденные ставки 🤔"},
			bot:  bot,
		},
	}
}

func (h *NackedBetsHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ставки, ожидающие подтверждения \n\nДинамо (Киев) - Шахтер (Донецк) 100$")
	// msg.ReplyMarkup = categoriesMenuKeyboard
	return msg
}
