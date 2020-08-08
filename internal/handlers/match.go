package handlers

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
)

type MatchHandler struct {
	GenericHandler
}

func NewMatchHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI) Handler {
	return &MatchHandler{
		GenericHandler{
			keys: []string{"match"},
			bot:  bot,
		},
	}
}

func (h *MatchHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	fmt.Println("%+v\n", update.Message.Text)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "История ставок \n\nДинамо (Киев) - Шахтер (Донецк) 100$ ❌\nУсик - Рокки 50$ ✅")
	return msg
}
