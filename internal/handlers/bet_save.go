package handlers

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/fireferretsbet/tg-bot/internal/orderbook"
	"gitlab.com/fireferretsbet/tg-bot/internal/serverenv"
)

var betSaveMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Главное меню ⬅️"),
	),
)

type BetSaveHandler struct {
	GenericHandler
}

func NewBetSaveHandler(env *serverenv.ServerEnv) Handler {
	return &BetSaveHandler{
		GenericHandler{
			keys: []string{"Сделать ставку 🤑"},
			env:  env,
		},
	}
}

func (h *BetSaveHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	state := h.env.UserManager().GetState(update.Message.From.ID)
	event := h.env.EventManager().Event(state.Event)
	user := h.env.UserManager().GetUser(update.Message.From.ID)

	order := orderbook.NewOrder()
	order.Side = state.Side
	order.Coeff = state.Coeff
	order.Qty = state.Qty

	_, err := event.PlaceOrder(user, order)
	if err != nil {
		if err == orderbook.ErrNotEnoughFounds {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Недостаточно средств 😞\n")
			msg.ReplyMarkup = balanceMenuKeyboard
			return msg
		}
		h.env.Logger().Error(err)
		return tgbotapi.NewMessage(update.Message.Chat.ID, "Не удалось сохранить ставку 😞\n")
	}

	text := fmt.Sprintf("Проверьте правильность данных\n\nМатч: %s\nСторона: %s\nКоэффициент: %s\nСтавка: %s$\n\n", event.Name, state.Side, state.Coeff, state.Qty)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = betCheckMenuKeyboard
	return msg
}
