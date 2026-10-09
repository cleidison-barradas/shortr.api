package utils

import (
	"github.com/cleidison-barradas/shortr.api/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type Hasher struct{}

var _ domain.PasswordHasher = (*Hasher)(nil)

func NewHasher() domain.PasswordHasher {
	return &Hasher{}
}

func (h *Hasher) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

func (h *Hasher) Compare(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
