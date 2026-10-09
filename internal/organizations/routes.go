package organizations

import (
	"net/http"

	"github.com/cleidison-barradas/shortr.api/internal/middleware"
)

func RegisterRoutes(mux *http.ServeMux, h *organiztionHandler, authMiddleware middleware.MiddlewareFunc) {
	mux.Handle("POST /api/organizations", authMiddleware(http.HandlerFunc(h.CreateHandler)))
	mux.Handle("GET /api/organizations", authMiddleware(http.HandlerFunc(h.GetOrganizationByUserHandler)))
}
