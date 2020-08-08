package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
	"gitlab.com/fireferretsbet/tg-bot/internal/user"
)

type BalanceHandler struct {
	GenericHandler
}

func NewBalanceHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI, userStates map[int]*user.UserState) Handler {
	return &BalanceHandler{
		GenericHandler{
			keys:       []string{"Баланс 🏦"},
			bot:        bot,
			userStates: userStates,
		},
	}
}

func (h *BalanceHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Баланс 🏦\n\nВыберите интересующий раздел.")
	// msg.ReplyMarkup = balanceMenuKeyboard
	return msg
}
