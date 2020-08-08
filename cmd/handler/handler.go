package handler

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
)

type HandlerFactory func(*logrus.Logger, *config.Config, *tgbotapi.BotAPI) Handler
type Handler interface {
	Handle(update tgbotapi.Update, ctx context.Context) tgbotapi.MessageConfig
	Keys() []string
}

type UpdateHandler struct {
	log      *logrus.Logger
	config   *config.Config
	bot      *tgbotapi.BotAPI
	handlers map[string]Handler
}

func NewHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI) *UpdateHandler {
	handlers := make(map[string]Handler)
	h := &UpdateHandler{
		log,
		config,
		bot,
		handlers,
	}

	h.RegisterHandlers([]HandlerFactory{
		NewStartHandler,
		NewCategoriesHandler,
		NewBalanceHandler,
		NewAckedBetsHandler,
		NewNackedBetsHandler,
		NewHistoryHandler,
		NewCategoryHandler,
	})

	return h
}

func (h *UpdateHandler) RegisterHandlers(factories []HandlerFactory) {
	for _, factory := range factories {
		handler := factory(h.log, h.config, h.bot)
		for _, key := range handler.Keys() {
			h.handlers[key] = handler
		}
	}
}

func (h *UpdateHandler) Handle(update tgbotapi.Update, ctx context.Context) {
	if update.Message == nil {
		return
	}

	h.log.WithFields(logrus.Fields{
		"user_id":  update.Message.From.ID,
		"username": update.Message.From.UserName,
		"text":     update.Message.Text,
	}).Info("message accepted")

	var msg tgbotapi.MessageConfig
	fallback := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, выберите пункт меню")

	if update.Message.IsCommand() {
		if handler, ok := h.handlers[update.Message.Command()]; ok {
			msg = handler.Handle(update, ctx)
		} else {
			msg = fallback
		}
	} else {
		if handler, ok := h.handlers[update.Message.Text]; ok {
			msg = handler.Handle(update, ctx)
		} else {
			msg = fallback
		}
	}

	h.bot.Send(msg)
}
