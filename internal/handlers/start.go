package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
	"gitlab.com/fireferretsbet/tg-bot/internal/user"
)

var startMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Категории 📂"),
		tgbotapi.NewKeyboardButton("Баланс 🏦"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Неподтвержденные ставки 🤔"),
		tgbotapi.NewKeyboardButton("Подтвержденные ставки ✅"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("История ставок 📜"),
	),
)

type StartHandler struct {
	GenericHandler
}

func NewStartHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI, userStates map[int]*user.UserState) Handler {
	return &StartHandler{
		GenericHandler{
			keys:       []string{"start", "Главное меню ⬅️"},
			bot:        bot,
			userStates: userStates,
		},
	}
}

func (h *StartHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вас приветствует FireFerrets бот для ставок. 🔥🐹🤑")
	msg.ReplyMarkup = startMenuKeyboard
	return msg
}
