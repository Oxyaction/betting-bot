package setup

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/fireferretsbet/tg-bot/internal/config"
	"gitlab.com/fireferretsbet/tg-bot/internal/event"
	"gitlab.com/fireferretsbet/tg-bot/internal/logger"
	"gitlab.com/fireferretsbet/tg-bot/internal/serverenv"
	"gitlab.com/fireferretsbet/tg-bot/internal/user"
)

func Setup(ctx context.Context) (*serverenv.ServerEnv, *config.Config, error) {
	config := config.NewConfig()
	log := logger.NewLogger(config)

	opts := []serverenv.Option{}

	bot, err := tgbotapi.NewBotAPI(config.ApiToken)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to connect to telegram bot api: %v", err)
	}
	opts = append(opts, serverenv.WithTelegramAPI(bot))

	em := event.NewManager(log)
	opts = append(opts, serverenv.WithEventManager(em))
	opts = append(opts, serverenv.WithLogger(log))
	opts = append(opts, serverenv.WithUserManager(user.NewUserManager()))

	return serverenv.New(ctx, opts...), config, nil
}
