package links

import (
	"net/http"

	"github.com/cleidison-barradas/shortr.api/internal/middleware"
)

func RegisterRoutes(mux *http.ServeMux, h *linksHandler, authMiddleware middleware.MiddlewareFunc) {
	mux.HandleFunc("GET /api/links/{code}", h.GetRedirectHandler)
	mux.Handle("GET /api/links/expand", authMiddleware(http.HandlerFunc(h.FindHandler)))
	mux.Handle("POST /api/links", authMiddleware(http.HandlerFunc(h.CreateHandler)))
}
