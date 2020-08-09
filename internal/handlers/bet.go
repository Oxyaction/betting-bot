package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/fireferretsbet/tg-bot/internal/serverenv"
	"gitlab.com/fireferretsbet/tg-bot/internal/utils"
)

var betMenuKeyboard = tgbotapi.NewReplyKeyboard(
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
