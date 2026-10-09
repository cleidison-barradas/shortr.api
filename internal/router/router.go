package router

import (
	"net/http"

	"github.com/cleidison-barradas/shortr.api/internal/config"
	"github.com/cleidison-barradas/shortr.api/internal/handler"
	"github.com/cleidison-barradas/shortr.api/internal/links"
	"github.com/cleidison-barradas/shortr.api/internal/middleware"
	"github.com/cleidison-barradas/shortr.api/internal/organizations"
	"github.com/cleidison-barradas/shortr.api/internal/signin"
	"github.com/cleidison-barradas/shortr.api/internal/users"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(pool *pgxpool.Pool, cfg config.Config) http.Handler {
	mux := http.NewServeMux()

	authMiddleware := middleware.AuthMiddleware(cfg.JwtSecret)

	mux.HandleFunc("GET /api/health", handler.Health)

	users.RegisterModule(mux, pool, authMiddleware)
	organizations.RegisterModule(mux, pool, authMiddleware)
	links.RegisterModule(mux, pool, authMiddleware)
	signin.RegisterModule(mux, pool, cfg)

	var h http.Handler = mux

	h = middleware.LoggingMiddleware(h)
	h = middleware.Recover(h)

	return h
}
