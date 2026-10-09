package domain

import (
	"errors"
	"time"
)

var (
	ErrInvalidOrganization  = errors.New("invalid organization")
	ErrOrganizationNotFound = errors.New("organization not found")
)

type Organization struct {
	id        string
	name      string
	enabled   bool
	userID    string
	siteURL   string
	role      string
	deletedAt *time.Time
	createdAt time.Time
	updatedAt *time.Time
}

type NewOrganizationParams struct {
	Name    string
	UserID  string
	SiteURL string
	Role    string
}

type RehydrateOrganizationParams struct {
	ID        string
	Name      string
	UserID    string
	Enabled   bool
	SiteURL   string
	Role      string
	DeletedAt *time.Time
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func NewOrganization(p *NewOrganizationParams) (*Organization, error) {

	if p.Name == "" || p.UserID == "" {
		return nil, ErrInvalidOrganization
	}

	return &Organization{
		name:      p.Name,
		userID:    p.UserID,
		enabled:   false,
		siteURL:   p.SiteURL,
		role:      p.Role,
		createdAt: time.Now().UTC(),
	}, nil
}

func RehydrateOrganization(p RehydrateOrganizationParams) (*Organization, error) {
	return &Organization{
		id:        p.ID,
		name:      p.Name,
		userID:    p.UserID,
		enabled:   p.Enabled,
		siteURL:   p.SiteURL,
		role:      p.Role,
		deletedAt: p.DeletedAt,
		createdAt: p.CreatedAt,
		updatedAt: p.UpdatedAt,
	}, nil
}

func (o *Organization) ID() string            { return o.id }
func (o *Organization) Name() string          { return o.name }
func (o *Organization) UserID() string        { return o.userID }
func (o *Organization) Enabled() bool         { return o.enabled }
func (o *Organization) SiteURL() string       { return o.siteURL }
func (o *Organization) Role() string          { return o.role }
func (o *Organization) DeletedAt() *time.Time { return o.deletedAt }
func (o *Organization) CreatedAt() time.Time  { return o.createdAt }
func (o *Organization) UpdatedAt() *time.Time { return o.updatedAt }
