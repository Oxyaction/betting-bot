package serverenv

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/event"
	"gitlab.com/fireferretsbet/tg-bot/internal/user"
)

type ServerEnv struct {
	eventManager *event.Manager
	bot          *tgbotapi.BotAPI
	log          *logrus.Logger
	um           *user.UserManager
}

// Option defines function types to modify the ServerEnv on creation.
type Option func(*ServerEnv) *ServerEnv

// New creates a new ServerEnv with the requested options.
func New(ctx context.Context, opts ...Option) *ServerEnv {
	env := &ServerEnv{}

	for _, f := range opts {
		env = f(env)
	}

	fmt.Printf("%#v\n", env)

	return env
}

func WithEventManager(em *event.Manager) Option {
	return func(s *ServerEnv) *ServerEnv {
		s.eventManager = em
		return s
	}
}

func WithTelegramAPI(bot *tgbotapi.BotAPI) Option {
	return func(s *ServerEnv) *ServerEnv {
		s.bot = bot
		return s
	}
}

func WithLogger(log *logrus.Logger) Option {
	return func(s *ServerEnv) *ServerEnv {
		s.log = log
		return s
	}
}

func WithUserManager(um *user.UserManager) Option {
	return func(s *ServerEnv) *ServerEnv {
		s.um = um
		return s
	}
}

func (e *ServerEnv) Bot() *tgbotapi.BotAPI {
	return e.bot
}

func (e *ServerEnv) EventManager() *event.Manager {
	return e.eventManager
}

func (e *ServerEnv) Logger() *logrus.Logger {
	return e.log
}

func (e *ServerEnv) UserManager() *user.UserManager {
	return e.um
}
