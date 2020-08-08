package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	ob "github.com/miktwon/orderbook"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
	"gitlab.com/fireferretsbet/tg-bot/internal/user"
)

var coeffMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1.2"),
		tgbotapi.NewKeyboardButton("1.5"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("5"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("10"),
		tgbotapi.NewKeyboardButton("20"),
		tgbotapi.NewKeyboardButton("Назад ⬅️"),
	),
)

type CoeffHandler struct {
	GenericHandler
}

func NewCoeffHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI, userStates map[int]*user.UserState) Handler {
	return &CoeffHandler{
		GenericHandler{
			keys:       []string{"coeff"},
			bot:        bot,
			userStates: userStates,
		},
	}
}

func (h *CoeffHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	text := update.Message.Text
	if text != "Назад ⬅️" {
		var side ob.Side
		if text == "Lay" {
			side = ob.Lay
		} else if text == "Back" {
			side = ob.Back
		} else {
			return tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите Lay или Back")
		}

		h.userStates[update.Message.From.ID].Side = side
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите коэффициент или введите свой")
	msg.ReplyMarkup = coeffMenuKeyboard
	return msg
}

func (h *CoeffHandler) GetDialogContext() string {
	return "bet"
}

func (h *CoeffHandler) GetPreviousRoute() string {
	return "side"
}
