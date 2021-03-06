package handlers

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/fireferretsbet/tg-bot/internal/serverenv"
	"gitlab.com/fireferretsbet/tg-bot/internal/utils"
)

var betCheckMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Сделать ставку 🤑"),
		tgbotapi.NewKeyboardButton("Назад ⬅️"),
	),
)

type BetCheckHandler struct {
	GenericHandler
}

func NewBetCheckHandler(env *serverenv.ServerEnv) Handler {
	return &BetCheckHandler{
		GenericHandler{
			keys: []string{"bet_check"},
			env:  env,
		},
	}
}

func (h *BetCheckHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	// persist bet to state
	if update.Message.Text != "Назад ⬅️" {
		bet, err := utils.DecimalFromText(update.Message.Text)
		if err != nil {
			return tgbotapi.NewMessage(update.Message.Chat.ID, "Неправильное значение")
		}
		h.env.UserManager().SetQty(update.Message.From.ID, bet)
	}
	state := h.env.UserManager().GetState(update.Message.From.ID)
	event := h.env.EventManager().Event(state.Event)

	text := fmt.Sprintf("Проверьте правильность данных\n\nМатч: *%s*\nСторона: *%s*\nКоэффициент: *%s*\nСтавка: *%s USDT*\n\n", event.Name, state.Side, state.Coeff, state.Qty)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = betCheckMenuKeyboard
	msg.ParseMode = "Markdown"
	return msg
}

func (h *BetCheckHandler) GetPreviousRoute() string {
	return "bet"
}
