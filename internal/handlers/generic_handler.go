package handlers

import "gitlab.com/fireferretsbet/tg-bot/internal/serverenv"

type GenericHandler struct {
	keys []string
	env  *serverenv.ServerEnv
}

func (h *GenericHandler) Keys() []string {
	return h.keys
}

func (h *GenericHandler) GetDialogContext() string {
	return ""
}

func (h *GenericHandler) GetPreviousRoute() string {
	return ""
}
