package organizations

import (
	"net/http"

	"github.com/cleidison-barradas/shortr.api/internal/infra/postgres"
	"github.com/cleidison-barradas/shortr.api/internal/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterModule(mux *http.ServeMux, pool *pgxpool.Pool, authMiddleware middleware.MiddlewareFunc) {

	orgRepo := postgres.NewOrganizationRepository(pool)
	userRepo := postgres.NewUserRepository(pool)

	srv := NewOrganizationService(userRepo, orgRepo)
	h := NewOrganizationHandler(srv)

	RegisterRoutes(mux, h, authMiddleware)
}
