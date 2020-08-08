package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
)

type AckedBetsHandler struct {
	GenericHandler
}

func NewAckedBetsHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI) Handler {
	return &AckedBetsHandler{
		GenericHandler{
			keys: []string{"Подтвержденные ставки ✅"},
			bot:  bot,
		},
	}
}

func (h *AckedBetsHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Подтвержденные ставки ✅\n\nДинамо (Киев) - Шахтер (Донецк) 100$")
	// msg.ReplyMarkup = categoriesMenuKeyboard
	return msg
}
