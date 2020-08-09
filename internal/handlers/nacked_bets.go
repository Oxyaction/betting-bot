package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/fireferretsbet/tg-bot/internal/serverenv"
)

type NackedBetsHandler struct {
	GenericHandler
}

func NewNackedBetsHandler(env *serverenv.ServerEnv) Handler {
	return &NackedBetsHandler{
		GenericHandler{
			keys: []string{"Неподтвержденные ставки 🤔"},
			env:  env,
		},
	}
}

func (h *NackedBetsHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ставки, ожидающие подтверждения \n\nДинамо (Киев) - Шахтер (Донецк) 100$")
	// msg.ReplyMarkup = categoriesMenuKeyboard
	return msg
}
