package signin

import (
	"net/http"

	"github.com/cleidison-barradas/shortr.api/internal/auth"
	"github.com/cleidison-barradas/shortr.api/internal/config"
	"github.com/cleidison-barradas/shortr.api/internal/infra/postgres"
	"github.com/cleidison-barradas/shortr.api/internal/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterModule(mux *http.ServeMux, pool *pgxpool.Pool, cfg config.Config) {
	userRepo := postgres.NewUserRepository(pool)

	hasher := utils.NewHasher()
	signInService := NewSiginService(userRepo, hasher)
	jwtService := auth.NewJWTService(cfg.JwtSecret)

	h := NewSigninHandler(signInService, jwtService)

	RegisterRoutes(mux, h)
}
