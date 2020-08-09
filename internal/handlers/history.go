package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/fireferretsbet/tg-bot/internal/serverenv"
)

type HistoryHandler struct {
	GenericHandler
}

func NewHistoryHandler(env *serverenv.ServerEnv) Handler {
	return &HistoryHandler{
		GenericHandler{
			keys: []string{"История ставок 📜"},
			env:  env,
		},
	}
}

func (h *HistoryHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "История ставок \n\nДинамо (Киев) - Шахтер (Донецк) *100 USDT* ❌\nУсик - Рокки *50 USDT* ✅")
	msg.ParseMode = "Markdown"
	return msg
}
