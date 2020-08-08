package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
	"gitlab.com/fireferretsbet/tg-bot/internal/user"
)

type NackedBetsHandler struct {
	GenericHandler
}

func NewNackedBetsHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI, userStates map[int]*user.UserState) Handler {
	return &NackedBetsHandler{
		GenericHandler{
			keys:       []string{"Неподтвержденные ставки 🤔"},
			bot:        bot,
			userStates: userStates,
		},
	}
}

func (h *NackedBetsHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ставки, ожидающие подтверждения \n\nДинамо (Киев) - Шахтер (Донецк) 100$")
	// msg.ReplyMarkup = categoriesMenuKeyboard
	return msg
}
