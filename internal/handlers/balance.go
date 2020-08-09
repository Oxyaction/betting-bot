package handlers

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/fireferretsbet/tg-bot/internal/serverenv"
	"gitlab.com/fireferretsbet/tg-bot/internal/utils"
)

var balanceMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Пополнить 💳"),
		tgbotapi.NewKeyboardButton("Вывести 💳"),
		tgbotapi.NewKeyboardButton("Главное меню ⬅️"),
	),
)

type BalanceHandler struct {
	GenericHandler
}

func NewBalanceHandler(env *serverenv.ServerEnv) Handler {
	return &BalanceHandler{
		GenericHandler{
			keys: []string{"Баланс 🏦", "top_up_success"},
			env:  env,
		},
	}
}

func (h *BalanceHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	user := h.env.UserManager().GetUser(update.Message.From.ID)
	var text string
	// top_up_success
	if h.env.UserManager().GetContextRoute(update.Message.From.ID) == "top_up_success" {
		incrBy, err := utils.DecimalFromText(update.Message.Text)
		if err != nil {
			return tgbotapi.NewMessage(update.Message.Chat.ID, "Неправильное значение")
		}

		user.ChangeBalance("top up", incrBy)
		text = "Баланс успешно пополнен ✅\n\n"
	} else {
		text = "Баланс 🏦\n\n"
	}
	text += fmt.Sprintf("Ваш текущий баланс: *%s USDT*.", user.GetBalance().Truncate(2).String())
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = balanceMenuKeyboard
	return msg
}
