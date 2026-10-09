package users

import (
	"net/http"

	"github.com/cleidison-barradas/shortr.api/internal/infra/postgres"
	"github.com/cleidison-barradas/shortr.api/internal/middleware"
	"github.com/cleidison-barradas/shortr.api/internal/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterModule(mux *http.ServeMux, pool *pgxpool.Pool, authMiddleware middleware.MiddlewareFunc) {
	repo := postgres.NewUserRepository(pool)
	hasher := utils.NewHasher()

	srv := NewUserService(repo, hasher)
	h := NewUserHandler(srv)

	RegisterRoutes(mux, h, authMiddleware)
}
