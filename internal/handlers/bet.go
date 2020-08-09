package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/fireferretsbet/tg-bot/internal/serverenv"
	"gitlab.com/fireferretsbet/tg-bot/internal/utils"
)

var betMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1USDT"),
		tgbotapi.NewKeyboardButton("2USDT"),
		tgbotapi.NewKeyboardButton("5USDT"),
		tgbotapi.NewKeyboardButton("10USDT"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("20USDT"),
		tgbotapi.NewKeyboardButton("50USDT"),
		tgbotapi.NewKeyboardButton("100USDT"),
		tgbotapi.NewKeyboardButton("Назад ⬅️"),
	),
)

type BetHandler struct {
	GenericHandler
}

func NewBetHandler(env *serverenv.ServerEnv) Handler {
	return &BetHandler{
		GenericHandler{
			keys: []string{"bet"},
			env:  env,
		},
	}
}

func (h *BetHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	if update.Message.Text != "Назад ⬅️" {
		coeff, err := utils.DecimalFromText(update.Message.Text)
		if err != nil {
			return tgbotapi.NewMessage(update.Message.Chat.ID, "Неправильное значение")
		}
		h.env.UserManager().SetCoeff(update.Message.From.ID, coeff)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Сделайте вашу ставку")
	msg.ReplyMarkup = betMenuKeyboard
	return msg
}

func (h *BetHandler) GetDialogContext() string {
	return "bet_check"
}

func (h *BetHandler) GetPreviousRoute() string {
	return "coeff"
}
