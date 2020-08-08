package logger

import (
	"os"

	"gitlab.com/fireferretsbet/tg-bot/internal/config"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func NewLogger(config *config.Config) *logrus.Logger {
	logger := logrus.New()
	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(os.Stdout)

	level, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		level = log.InfoLevel
	}

	logger.SetLevel(level)

	return logger
}
