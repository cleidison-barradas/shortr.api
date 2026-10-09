package users

import (
	"net/http"

	"github.com/cleidison-barradas/shortr.api/internal/middleware"
)

func RegisterRoutes(mux *http.ServeMux, h *UserHandler, authMiddleware middleware.MiddlewareFunc) {
	mux.HandleFunc("POST /api/users", h.CreateHandler)
	mux.Handle("GET /api/users", authMiddleware(http.HandlerFunc(h.GetHandler)))
	mux.Handle("PUT /api/users/{userID}", authMiddleware(http.HandlerFunc(h.UpdateHandler)))
}
