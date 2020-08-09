package main

import (
	"context"
	"net/http"
	"strconv"

	"gitlab.com/fireferretsbet/tg-bot/internal/handlers"
	"gitlab.com/fireferretsbet/tg-bot/internal/setup"
)

func main() {
	ctx := context.Background()
	env, config, err := setup.Setup(ctx)
	if err != nil {
		panic(err)
	}
	updates := env.Bot().ListenForWebhook("/")
	go http.ListenAndServe("0.0.0.0:"+strconv.Itoa(config.Port), nil)

	h := handlers.NewUpdateHandler(env)
	for update := range updates {
		go h.Handle(update, ctx)
	}
}
