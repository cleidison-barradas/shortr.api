package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	SecretKey string
}

type GenerateTokenParams struct {
	UserID string
	OrgID  string
	Email  string
}

func NewJWTService(secretKey string) *JWTService {
	return &JWTService{
		SecretKey: secretKey,
	}
}

func (s *JWTService) GenerateToken(p *GenerateTokenParams) (string, error) {
	claims := jwt.MapClaims{
		"user_id": p.UserID,
		"org_id":  p.OrgID,
		"email":   p.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.SecretKey))
}

func (s *JWTService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.SecretKey), nil
	})
}
