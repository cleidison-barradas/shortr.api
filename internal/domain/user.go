package domain

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrInvalidUser            = errors.New("invalid user")
	ErrUserNotFound           = errors.New("user not found")
	ErrUserEmailAlreadyExists = errors.New("user email already exists")
)

type User struct {
	id           string
	firstName    string
	lastName     string
	email        Email
	passwordHash string
	lastLogin    *time.Time
	createdAt    time.Time
	updatedAt    *time.Time
}

type NewUserParams struct {
	FirstName    string
	LastName     string
	Email        Email
	PasswordHash string
}

type PasswordHasher interface {
	Compare(hashedPassword, password string) bool
	HashPassword(password string) (string, error)
}

func NewUser(p NewUserParams) (*User, error) {

	if p.FirstName == "" || p.LastName == "" || p.PasswordHash == "" {
		return nil, fmt.Errorf("%w: fistname,lastname,password are required", ErrInvalidUser)
	}

	if p.Email == (Email{}) {
		return nil, fmt.Errorf("%w: email is required", ErrInvalidUser)
	}

	now := time.Now().UTC()

	return &User{
		firstName:    p.FirstName,
		lastName:     p.LastName,
		email:        p.Email,
		passwordHash: p.PasswordHash,
		lastLogin:    nil,
		createdAt:    now,
		updatedAt:    nil,
	}, nil
}

type RehydratedUserParams struct {
	ID           string
	FirstName    string
	LastName     string
	Email        Email
	PasswordHash string
	LastLogin    *time.Time
	CreatedAt    time.Time
	UpdatedAt    *time.Time
}

func RehydrateUser(p RehydratedUserParams) (*User, error) {
	return &User{
		id:           p.ID,
		firstName:    p.FirstName,
		lastName:     p.LastName,
		email:        p.Email,
		passwordHash: p.PasswordHash,
		lastLogin:    p.LastLogin,
		createdAt:    p.CreatedAt,
		updatedAt:    p.UpdatedAt,
	}, nil
}

func (u *User) CheckPassword(password string, h PasswordHasher) bool {
	return h.Compare(u.passwordHash, password)
}

func (u *User) HashPassword(password string, h PasswordHasher) (string, error) {
	return h.HashPassword(password)
}

func (u *User) ID() string            { return u.id }
func (u *User) FirstName() string     { return u.firstName }
func (u *User) LastName() string      { return u.lastName }
func (u *User) FullName() string      { return u.firstName + " " + u.lastName }
func (u *User) Email() Email          { return u.email }
func (u *User) PasswordHash() string  { return u.passwordHash }
func (u *User) LastLogin() *time.Time { return u.lastLogin }
func (u *User) CreatedAt() time.Time  { return u.createdAt }
func (u *User) UpdatedAt() *time.Time { return u.updatedAt }
