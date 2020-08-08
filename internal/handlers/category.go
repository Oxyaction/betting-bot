package handlers

import (
	"context"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
)

var matches map[string][]string = map[string][]string{
	"Спорт ⚽":      []string{"1. Динамо - Шахтер", "2. Усик - Рокки"},
	"Киберспорт 🎮": []string{"1. Omega - Spirits (Dota 2)", "2. Scarlett - Neeb (SC2)", "3. Moon - Grubby (WC3)"},
	"Политика 🏛️":  []string{"1. Joe Biden - Donald Trump", "2. Лукашенко - Тихановская"},
}

var categoryMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Спорт ⚽"),
		tgbotapi.NewKeyboardButton("Киберспорт 🎮"),
		tgbotapi.NewKeyboardButton("Политика 🏛️"),
		tgbotapi.NewKeyboardButton("Главное меню ⬅️"),
	),
)

type CategoryHandler struct {
	GenericHandler
}

func NewCategoryHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI) Handler {
	return &CategoryHandler{
		GenericHandler{
			keys: []string{
				"Спорт ⚽",
				"Киберспорт 🎮",
				"Политика 🏛️",
			},
			bot: bot,
		},
	}
}

func (h *CategoryHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	var textResponse string
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	if matches, ok := matches[update.Message.Text]; ok {
		var buttons []tgbotapi.KeyboardButton = make([]tgbotapi.KeyboardButton, len(matches))
		for i := 0; i < len(matches); i++ {
			buttons = append(buttons, tgbotapi.NewKeyboardButton(matches[i]))
		}
		buttons = append(buttons, tgbotapi.NewKeyboardButton("Главное меню ⬅️"))
		var digitsMenuKeyboard = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(buttons...),
		)

		textResponse = "Выберите интересующий матч: \n\n"
		textResponse += strings.Join(matches, "\n")
		msg.Text = textResponse
		msg.ReplyMarkup = digitsMenuKeyboard
	} else {
		msg.Text = "Выберите категорию из меню"
	}

	return msg
}

func (h *CategoryHandler) GetDialogContext() string {
	return "match"
}
