package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
	"gitlab.com/fireferretsbet/tg-bot/internal/user"
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

func NewBetHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI, userStates map[int]*user.UserState) Handler {
	return &BetHandler{
		GenericHandler{
			keys:       []string{"bet"},
			bot:        bot,
			userStates: userStates,
		},
	}
}

func (h *BetHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	if update.Message.Text != "Назад ⬅️" {
		coeff, err := utils.DecimalFromText(update.Message.Text)
		if err != nil {
			return tgbotapi.NewMessage(update.Message.Chat.ID, "Неправильное значение")
		}

		h.userStates[update.Message.From.ID].Coeff = coeff
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
