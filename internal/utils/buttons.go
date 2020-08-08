package utils

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func BuildKeyboardFromStrings(values []string) tgbotapi.ReplyKeyboardMarkup {
	// var buttons []tgbotapi.KeyboardButton
	buttons := make([]tgbotapi.KeyboardButton, len(values))
	for _, category := range values {
		buttons = append(buttons, tgbotapi.NewKeyboardButton(category))
	}

	return tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(buttons...))
}
