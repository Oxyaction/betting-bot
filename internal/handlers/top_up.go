package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
	"gitlab.com/fireferretsbet/tg-bot/internal/user"
)

var topUpMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1$"),
		tgbotapi.NewKeyboardButton("2$"),
		tgbotapi.NewKeyboardButton("5$"),
		tgbotapi.NewKeyboardButton("10$"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("20$"),
		tgbotapi.NewKeyboardButton("50$"),
		tgbotapi.NewKeyboardButton("100$"),
		tgbotapi.NewKeyboardButton("Назад ⬅️"),
	),
)

type TopUpHandler struct {
	GenericHandler
}

func NewTopUpHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI, userStates map[int]*user.UserState) Handler {
	return &TopUpHandler{
		GenericHandler{
			keys:       []string{"Пополнить 💳"},
			bot:        bot,
			userStates: userStates,
		},
	}
}

func (h *TopUpHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите сумму пополнения")
	msg.ReplyMarkup = topUpMenuKeyboard
	return msg
}

func (h *TopUpHandler) GetDialogContext() string {
	return "top_up_success"
}

func (h *TopUpHandler) GetPreviousRoute() string {
	return "balance"
}
