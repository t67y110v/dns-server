package app

import (
	"log/slog"
	"net/http"

	"github.com/t67y110v/dns-server/internal/config"
	"github.com/t67y110v/dns-server/internal/handlers"
)

type Server struct {
	config   *config.Config
	logger   *slog.Logger
	handlers *handlers.Handlers
}

func New(config *config.Config, logger *slog.Logger) *Server {
	handlers := handlers.NewHandlers(config, logger)
	return &Server{
		config:   config,
		logger:   logger,
		handlers: handlers,
	}
}

func (s *Server) setRoutes() {
	http.HandleFunc("/endpoints", s.handlers.GetEndpoints)
	http.HandleFunc("/status", s.handlers.Status)
}

func (s *Server) Start() error {
	s.setRoutes()
	s.logger.Info("Starting application")
	return http.ListenAndServe(s.config.BasePath+s.config.BindAddr, nil)

}
