package domain

import (
	"errors"
	"regexp"
	"strings"
)

var ErrInvalidEmail = errors.New("invalid email")

type Email struct {
	value string
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")

func NewEmail(raw string) (Email, error) {
	raw = strings.TrimSpace(strings.ToLower(raw))
	if raw == "" || !emailRegex.MatchString(raw) {
		return Email{}, ErrInvalidEmail
	}

	return Email{value: raw}, nil
}

func (e Email) String() string { return e.value }
