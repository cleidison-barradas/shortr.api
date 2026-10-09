package domain

import (
	"errors"
	"time"
)

var (
	ErrInvalidLink           = errors.New("invalid link")
	ErrLinkNotFound          = errors.New("link not found")
	ErrShorlinkAlreadyExists = errors.New("short link already exists")
)

type Link struct {
	id              string
	short_link      string
	original_link   string
	organization_id string
	alias           string
	expires_in      *time.Time
	created_at      time.Time
	updated_at      *time.Time
}

type NewLinkParams struct {
	ShortURL       string
	Alias          string
	OriginalURL    string
	ExpiresIn      *time.Time
	OrganizationID string
}

func NewLink(p NewLinkParams) (*Link, error) {
	if p.ShortURL == "" || p.OriginalURL == "" {
		return nil, ErrInvalidLink
	}

	return &Link{
		short_link:      p.ShortURL,
		original_link:   p.OriginalURL,
		alias:           p.Alias,
		expires_in:      p.ExpiresIn,
		organization_id: p.OrganizationID,
		created_at:      time.Now(),
	}, nil
}

type RehydrateLinkParams struct {
	ID             string
	ShortURL       string
	OriginalURL    string
	OrganizationID string
	Alias          string
	ExpiresIn      *time.Time
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}

func RehydrateLink(p RehydrateLinkParams) (*Link, error) {
	return &Link{
		id:              p.ID,
		short_link:      p.ShortURL,
		original_link:   p.OriginalURL,
		organization_id: p.OrganizationID,
		alias:           p.Alias,
		expires_in:      p.ExpiresIn,
		created_at:      p.CreatedAt,
		updated_at:      p.UpdatedAt,
	}, nil
}

func (l *Link) IsExpired() bool {
	now := time.Now().UTC()

	if l.expires_in != nil && l.expires_in.Before(now) {
		return true
	}

	return false
}

func (l *Link) ID() string             { return l.id }
func (l *Link) ShortURL() string       { return l.short_link }
func (l *Link) OriginalURL() string    { return l.original_link }
func (l *Link) OrganizationID() string { return l.organization_id }
func (l *Link) Alias() string          { return l.alias }
func (l *Link) ExpiresIn() *time.Time  { return l.expires_in }
func (l *Link) CreatedAt() time.Time   { return l.created_at }
func (l *Link) UpdatedAt() *time.Time  { return l.updated_at }
