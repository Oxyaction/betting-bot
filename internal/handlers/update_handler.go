package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/serverenv"
)

type HandlerFactory func(*serverenv.ServerEnv) Handler

type UpdateHandler struct {
	env      *serverenv.ServerEnv
	handlers map[string]Handler
	// userStates map[int]*user.UserState
}

func NewUpdateHandler(env *serverenv.ServerEnv) *UpdateHandler {
	handlers := make(map[string]Handler)
	// userStates := make(map[int]*user.UserState)
	h := &UpdateHandler{
		env,
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
		NewCoeffHandler,
		NewBetHandler,
		NewBetCheckHandler,
		NewTopUpHandler,
		NewSideHandler,
		NewBetSaveHandler,
	})

	return h
}

func (h *UpdateHandler) RegisterHandlers(factories []HandlerFactory) {
	for _, factory := range factories {
		handler := factory(h.env)
		for _, key := range handler.Keys() {
			h.handlers[key] = handler
		}
	}
}

func (h *UpdateHandler) Handle(update tgbotapi.Update, ctx context.Context) error {
	if update.Message == nil {
		return nil
	}

	userId := update.Message.From.ID
	userManager := h.env.UserManager()

	h.env.Logger().WithFields(logrus.Fields{
		"user_id":  userId,
		"username": update.Message.From.UserName,
		"text":     update.Message.Text,
	}).Info("message accepted")

	userManager.Add(update.Message.From.ID)

	var msg tgbotapi.MessageConfig
	fallback := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная комманда")

	var handler Handler
	var route string

	// handle back btn click
	if update.Message.Text == "Назад ⬅️" {
		route = userManager.GetPreviousRoute(userId)
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
		contextRoute := userManager.GetContextRoute(userId)
		if contextRoute != "" {
			handler = h.handlers[contextRoute]
		}
	}

	if handler != nil {
		msg = handler.Handle(update, ctx)
		err := userManager.SetContextRoute(userId, handler.GetDialogContext())
		if err != nil {
			return err
		}
		err = userManager.SetPreviousRoute(userId, handler.GetPreviousRoute())
		if err != nil {
			return err
		}
	} else {
		msg = fallback
	}

	// fmt.Printf("%#v\n", h.userStates[update.Message.From.ID])

	// h.bot.Send(msg)
	h.env.Bot().Send(msg)
	return nil
}
