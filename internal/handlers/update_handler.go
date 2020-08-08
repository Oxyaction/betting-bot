package handlers

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
	"gitlab.com/fireferretsbet/tg-bot/internal/user"
)

type HandlerFactory func(*logrus.Logger, *config.Config, *tgbotapi.BotAPI, map[int]*user.UserState) Handler

type UpdateHandler struct {
	log        *logrus.Logger
	config     *config.Config
	bot        *tgbotapi.BotAPI
	handlers   map[string]Handler
	userStates map[int]*user.UserState
}

func NewUpdateHandler(log *logrus.Logger, config *config.Config, bot *tgbotapi.BotAPI) *UpdateHandler {
	handlers := make(map[string]Handler)
	userStates := make(map[int]*user.UserState)
	h := &UpdateHandler{
		log,
		config,
		bot,
		handlers,
		userStates,
	}

	h.RegisterHandlers([]HandlerFactory{
		NewStartHandler,
		NewCategoriesHandler,
		NewBalanceHandler,
		NewAckedBetsHandler,
		NewNackedBetsHandler,
		NewHistoryHandler,
		NewCategoryHandler,
		NewCoeffHandler,
		NewBetHandler,
		NewBetCheckHandler,
		NewTopUpHandler,
		NewSideHandler,
	})

	return h
}

func (h *UpdateHandler) RegisterHandlers(factories []HandlerFactory) {
	for _, factory := range factories {
		handler := factory(h.log, h.config, h.bot, h.userStates)
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

	if _, ok := h.userStates[update.Message.From.ID]; !ok {
		h.userStates[update.Message.From.ID] = &user.UserState{
			PreviousRoute: "start",
		}
	}

	var msg tgbotapi.MessageConfig
	fallback := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная комманда")

	var handler Handler
	var route string

	// handle back btn click
	if update.Message.Text == "Назад ⬅️" {
		fmt.Println("xxx")
		fmt.Printf("%#v\n", h.userStates[update.Message.From.ID])
		route = h.userStates[update.Message.From.ID].PreviousRoute
	} else {
		route = update.Message.Text
	}

	if update.Message.IsCommand() {
		handler = h.handlers[update.Message.Command()]
		// trying to get text handler
	} else if _, ok := h.handlers[route]; ok {
		handler = h.handlers[route]
		// trying to get contextual handler
	} else {
		if h.userStates[update.Message.From.ID].ContextRoute != "" {
			handler = h.handlers[h.userStates[update.Message.From.ID].ContextRoute]
		}
	}

	if handler != nil {
		msg = handler.Handle(update, ctx)
		h.userStates[update.Message.From.ID].ContextRoute = handler.GetDialogContext()
		h.userStates[update.Message.From.ID].PreviousRoute = handler.GetPreviousRoute()
	} else {
		msg = fallback
	}

	fmt.Printf("%#v\n", h.userStates[update.Message.From.ID])

	h.bot.Send(msg)
}
