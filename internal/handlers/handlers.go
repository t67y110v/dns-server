package handlers

import (
	"log/slog"

	"github.com/t67y110v/dns-server/internal/config"
)

type Handlers struct {
	logger *slog.Logger
	config *config.Config
}

func NewHandlers(config *config.Config, logger *slog.Logger) *Handlers {
	return &Handlers{
		logger: logger,
		config: config,
	}
}
