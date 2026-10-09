package app

import (
	"context"
	"errors"

	"github.com/cleidison-barradas/shortr.api/internal/domain"
)

var (
	ErrBodyRequest     = errors.New("app: invalid body request")
	ErrMissingUserID   = errors.New("app: missing user id")
	ErrMissingOrgID    = errors.New("app: missing org id")
	ErrInvalidPassword = errors.New("app: invalid password")
	ErrMissingURLCode  = errors.New("app: missing url code")
	ErrLinkExpired     = errors.New("app: link expired")
)

type UserRepository interface {
	Find(ctx context.Context) ([]domain.User, error)
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByID(ctx context.Context, id string) (*domain.User, error)
}

type OrganizationRepository interface {
	FindByUserID(ctx context.Context, orgID string) (*domain.Organization, error)
	Create(ctx context.Context, org *domain.Organization) (*domain.Organization, error)
}

type LinkRepository interface {
	FindByShortLink(ctx context.Context, shortLink string) (*domain.Link, error)
	Create(ctx context.Context, link *domain.Link) (*domain.Link, error)
	Update(ctx context.Context, link *domain.Link) (*domain.Link, error)
	FindByID(ctx context.Context, id string) (*domain.Link, error)
}
