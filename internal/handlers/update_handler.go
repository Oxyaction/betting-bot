package handlers

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
)

type HandlerFactory func(*logrus.Logger, *config.Config, *tgbotapi.BotAPI) Handler

type UpdateHandler struct {
	log         *logrus.Logger
	config      *config.Config
	bot         *tgbotapi.BotAPI
	handlers    map[string]Handler
	userContext map[int]string
}

func NewUpdateHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI) *UpdateHandler {
	handlers := make(map[string]Handler)
	userContext := make(map[int]string)
	h := &UpdateHandler{
		log,
		config,
		bot,
		handlers,
		userContext,
	}

	h.RegisterHandlers([]HandlerFactory{
		NewStartHandler,
		NewCategoriesHandler,
		NewBalanceHandler,
		NewAckedBetsHandler,
		NewNackedBetsHandler,
		NewHistoryHandler,
		NewCategoryHandler,
		NewMatchHandler,
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
	fallback := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная комманда")

	var handler Handler
	// trying to get command handler
	if update.Message.IsCommand() {
		handler = h.handlers[update.Message.Command()]
		// trying to get text handler
	} else if _, ok := h.handlers[update.Message.Text]; ok {
		handler = h.handlers[update.Message.Text]
		// trying to get contextual handler
	} else if _, ok := h.userContext[update.Message.From.ID]; ok {
		key := h.userContext[update.Message.From.ID]
		if key != "" {
			handler = h.handlers[key]
		}
	}

	if handler != nil {
		msg = handler.Handle(update, ctx)
		h.userContext[update.Message.From.ID] = handler.GetDialogContext()
	} else {
		msg = fallback
	}

	fmt.Printf("%+v\n", h.userContext)

	h.bot.Send(msg)
}
