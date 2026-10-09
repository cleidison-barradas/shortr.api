package signin

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *SigninHandler) {
	mux.HandleFunc("POST /api/signin", h.SignInHandler)
}