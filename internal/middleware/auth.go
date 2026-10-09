package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/cleidison-barradas/shortr.api/internal/auth"
	"github.com/cleidison-barradas/shortr.api/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

type Role string

var (
	ErrMissingAuthHeader = errors.New("missing auth header")
	ErrInvalidJWTToken   = errors.New("invalid jwt token")
)

const (
	UserRole  Role = "user"
	AdminRole Role = "admin"
)

type AuthProps struct {
	OrgID  string
	UserID string
	Role   Role
}

func AuthMiddleware(secret string) func(http.Handler) http.Handler {
	srv := auth.NewJWTService(secret)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				utils.Error(w, r, ErrMissingAuthHeader)
				return
			}

			headerToken := strings.TrimPrefix(authHeader, "Bearer ")
			accessToken, err := srv.ValidateToken(headerToken)

			if err != nil || !accessToken.Valid {
				utils.Error(w, r, ErrInvalidJWTToken)
				return
			}

			claims, ok := accessToken.Claims.(jwt.MapClaims)

			if !ok {
				utils.Error(w, r, ErrInvalidJWTToken)
				return
			}

			userID := claims["user_id"].(string)

			ctx := context.WithValue(r.Context(), "user_id", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
