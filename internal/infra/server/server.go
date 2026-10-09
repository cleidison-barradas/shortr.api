package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/cleidison-barradas/shortr.api/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func New(cfg config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         ":" + cfg.Port,
			Handler:      handler,
			ReadTimeout:  cfg.ReadTimeout,
			IdleTimeout:  cfg.IdleTimeout,
			WriteTimeout: cfg.WriteTimeout,
		},
	}
}

func (s *Server) Listen() error {
	slog.Info("Http server running", "Addr", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
