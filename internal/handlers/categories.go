package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/fireferretsbet/tg-bot/internal/serverenv"
	"gitlab.com/fireferretsbet/tg-bot/internal/utils"
)

var categoriesMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Спорт ⚽"),
		tgbotapi.NewKeyboardButton("Киберспорт 🎮"),
		tgbotapi.NewKeyboardButton("Политика 🏛️"),
		tgbotapi.NewKeyboardButton("Главное меню ⬅️"),
	),
)

type CategoriesHandler struct {
	GenericHandler
}

func NewCategoriesHandler(env *serverenv.ServerEnv) Handler {
	return &CategoriesHandler{
		GenericHandler{
			keys: []string{"Категории 📂", "categories"},
			env:  env,
		},
	}
}

func (h *CategoriesHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Чтобы сделать ставку выберите интересующую категорию. 📂")
	categories := h.env.EventManager().Categories()
	categories = append(categories, "Главное меню ⬅️")
	msg.ReplyMarkup = utils.BuildKeyboardFromStrings(categories)
	return msg
}

func (h *CategoriesHandler) GetDialogContext() string {
	return "category"
}
