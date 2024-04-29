package app

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/t67y110v/dns-server/internal/config"
	"github.com/t67y110v/dns-server/internal/handlers"
)

type Server struct {
	config   *config.Config
	logger   *slog.Logger
	handlers *handlers.Handlers
	router   *chi.Mux
}

func New(config *config.Config, logger *slog.Logger) *Server {
	handlers := handlers.NewHandlers(logger)
	return &Server{
		router:   chi.NewRouter(),
		config:   config,
		logger:   logger,
		handlers: handlers,
	}
}

func (s *Server) setRoutes() {
	s.router.Use(middleware.Logger)
	s.router.Route("/endpoints", func(r chi.Router) {
		r.Route("/{target}", func(r chi.Router) {
			r.Get("/", s.handlers.GetEndpoints)
			r.Post("/", s.handlers.PostEndpoints)
			r.Patch("/", s.handlers.PatchEndpoints)
		})
	})

	s.router.Get("/status", s.handlers.GetStatus)
}

func (s *Server) Start() error {

	s.setRoutes()
	s.logger.Info("Starting application")

	return http.ListenAndServe(s.config.BasePath+s.config.BindAddr, s.router)

}
