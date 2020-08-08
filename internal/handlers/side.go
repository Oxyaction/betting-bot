package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
	"gitlab.com/fireferretsbet/tg-bot/internal/user"
)

var sideMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Back"),
		tgbotapi.NewKeyboardButton("Lay"),
	),
)

type SideHandler struct {
	GenericHandler
}

func NewSideHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI, userStates map[int]*user.UserState) Handler {
	return &SideHandler{
		GenericHandler{
			keys:       []string{"side"},
			bot:        bot,
			userStates: userStates,
		},
	}
}

func (h *SideHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	h.userStates[update.Message.From.ID].Match = update.Message.Text

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите сторону")
	msg.ReplyMarkup = sideMenuKeyboard
	return msg
}

func (h *SideHandler) GetDialogContext() string {
	return "coeff"
}

func (h *SideHandler) GetPreviousRoute() string {
	return "categories"
}
