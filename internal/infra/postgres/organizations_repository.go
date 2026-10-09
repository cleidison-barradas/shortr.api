package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/cleidison-barradas/shortr.api/internal/app"
	"github.com/cleidison-barradas/shortr.api/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type organizationRepository struct {
	pool *pgxpool.Pool
}

func NewOrganizationRepository(pool *pgxpool.Pool) app.OrganizationRepository {
	return &organizationRepository{
		pool: pool,
	}
}

func (r *organizationRepository) Create(ctx context.Context, org *domain.Organization) (*domain.Organization, error) {
	exec := executor(ctx, r.pool)
	query := `
		INSERT INTO organizations (name, user_id, enabled, site_url, role, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	var (
		id string
	)

	err := exec.QueryRow(
		ctx,
		query,
		org.Name(),
		org.UserID(),
		org.Enabled(),
		org.SiteURL(),
		org.Role(),
		org.CreatedAt(),
	).Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("postgres create: %w", err)
	}

	return domain.RehydrateOrganization(domain.RehydrateOrganizationParams{
		ID:        id,
		Name:      org.Name(),
		UserID:    org.ID(),
		Enabled:   org.Enabled(),
		SiteURL:   org.SiteURL(),
		Role:      org.Role(),
		CreatedAt: org.CreatedAt(),
	})
}

func (r *organizationRepository) FindByUserID(ctx context.Context, userID string) (*domain.Organization, error) {
	exec := executor(ctx, r.pool)
	query := `
		SELECT
			id,
			name,
			user_id,
			enabled,
			site_url,
			role,
			created_at,
			updated_at
		FROM organizations
		WHERE user_id = $1
	`
	var (
		id        string
		name      string
		user_id   string
		enabled   bool
		site_url  string
		role      string
		createdAt time.Time
		updatedAt *time.Time
	)

	err := exec.QueryRow(
		ctx,
		query,
		userID,
	).Scan(
		&id,
		&name,
		&user_id,
		&enabled,
		&site_url,
		&role,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrOrganizationNotFound
		}
		return nil, fmt.Errorf("postgres find by user id: %w", err)
	}

	return domain.RehydrateOrganization(domain.RehydrateOrganizationParams{
		ID:        id,
		Name:      name,
		UserID:    userID,
		Enabled:   enabled,
		SiteURL:   site_url,
		Role:      role,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
}
