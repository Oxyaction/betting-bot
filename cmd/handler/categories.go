package handler

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
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
	keys []string
	bot  *tgbotapi.BotAPI
}

func NewCategoriesHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI) Handler {
	return &CategoriesHandler{
		keys: []string{"Категории 📂"},
		bot:  bot,
	}
}

func (h *CategoriesHandler) Keys() []string {
	return h.keys
}

func (h *CategoriesHandler) Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Чтобы сделать ставку выберите интересующую категорию. 📂")
	msg.ReplyMarkup = categoriesMenuKeyboard
	return msg
}
