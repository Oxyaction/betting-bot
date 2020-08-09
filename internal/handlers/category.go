package handlers

import (
	"context"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/fireferretsbet/tg-bot/internal/serverenv"
	"gitlab.com/fireferretsbet/tg-bot/internal/utils"
)

// var matches map[string][]string = map[string][]string{
// 	"Спорт ⚽":      []string{"1. Динамо - Шахтер", "2. Усик - Рокки"},
// 	"Киберспорт 🎮": []string{"1. Omega - Spirits (Dota 2)", "2. Scarlett - Neeb (SC2)", "3. Moon - Grubby (WC3)"},
// 	"Политика 🏛️":  []string{"1. Joe Biden - Donald Trump", "2. Лукашенко - Тихановская"},
// }

type CategoryHandler struct {
	GenericHandler
}

func NewCategoryHandler(env *serverenv.ServerEnv) Handler {
	return &CategoryHandler{
		GenericHandler{
			keys: []string{
				"category",
			},
			env: env,
		},
	}
}

func (h *CategoryHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	var textResponse string
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	events := h.env.EventManager().EventsByCategory(update.Message.Text)
	eventTitles := make([]string, len(events))
	for _, event := range events {
		eventTitles = append(eventTitles, event.Name)
	}

	textResponse = "Выберите интересующий матч:\n"
	textResponse += strings.Join(eventTitles, "\n")
	msg.Text = textResponse

	eventTitles = append(eventTitles, "Назад ⬅️")
	keyboard := utils.BuildKeyboardFromStrings(eventTitles)
	msg.ReplyMarkup = keyboard

	return msg
}

func (h *CategoryHandler) GetDialogContext() string {
	return "side"
}

func (h *CategoryHandler) GetPreviousRoute() string {
	return "categories"
}
