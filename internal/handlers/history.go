package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
	"gitlab.com/fireferretsbet/tg-bot/internal/user"
)

type HistoryHandler struct {
	GenericHandler
}

func NewHistoryHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI, userStates map[int]*user.UserState) Handler {
	return &HistoryHandler{
		GenericHandler{
			keys:       []string{"История ставок 📜"},
			bot:        bot,
			userStates: userStates,
		},
	}
}

func (h *HistoryHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "История ставок \n\nДинамо (Киев) - Шахтер (Донецк) 100$ ❌\nУсик - Рокки 50$ ✅")
	// msg.ReplyMarkup = categoriesMenuKeyboard
	return msg
}
