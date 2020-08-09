package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	ob "github.com/miktwon/orderbook"
	"gitlab.com/fireferretsbet/tg-bot/internal/serverenv"
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

func NewCoeffHandler(env *serverenv.ServerEnv) Handler {
	return &CoeffHandler{
		GenericHandler{
			keys: []string{"coeff"},
			env:  env,
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

		h.env.UserManager().SetSide(update.Message.From.ID, side)
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
