package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/fireferretsbet/tg-bot/internal/serverenv"
)

type AckedBetsHandler struct {
	GenericHandler
}

func NewAckedBetsHandler(env *serverenv.ServerEnv) Handler {
	return &AckedBetsHandler{
		GenericHandler{
			keys: []string{"Подтвержденные ставки ✅"},
			env:  env,
		},
	}
}

func (h *AckedBetsHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Подтвержденные ставки ✅\n\nДинамо (Киев) - Шахтер (Донецк) *100 USDT*")
	msg.ParseMode = "Markdown"
	return msg
}
