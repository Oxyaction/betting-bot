package controllers

import (
	"fmt"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/fireferretsbet/tg-bot/internal/serverenv"
)

type SettleController struct {
	env *serverenv.ServerEnv
}

func NewSettleController(env *serverenv.ServerEnv) *SettleController {
	return &SettleController{env}
}

func (sc *SettleController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	text := "*Поздравляем* 🎉\n⚽ Матч *Шахтер - Динамо* завершился\nВы выиграли *76 USDT* 💵"
	// hardcoded oxyaction
	msg := tgbotapi.NewMessage(105040780, text)
	msg.ParseMode = "Markdown"

	sc.env.Bot().Send(msg)
	fmt.Fprintf(w, "Settled")
}
