package links

import (
	"net/http"

	"github.com/cleidison-barradas/shortr.api/internal/infra/postgres"
	"github.com/cleidison-barradas/shortr.api/internal/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterModule(mux *http.ServeMux, pool *pgxpool.Pool, authMiddleware middleware.MiddlewareFunc) {
	linkRepo := postgres.NewLinkRepository(pool)
	srv := newLinksService(linkRepo)
	h := newLinksHandler(srv)

	RegisterRoutes(mux, h, authMiddleware)
}
