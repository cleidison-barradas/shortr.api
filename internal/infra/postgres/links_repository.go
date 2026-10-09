package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cleidison-barradas/shortr.api/internal/app"
	"github.com/cleidison-barradas/shortr.api/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type linkRepository struct {
	pool *pgxpool.Pool
}

func NewLinkRepository(pool *pgxpool.Pool) app.LinkRepository {
	return &linkRepository{
		pool: pool,
	}
}

func (r *linkRepository) Create(ctx context.Context, link *domain.Link) (*domain.Link, error) {
	exec := executor(ctx, r.pool)
	query := `
		INSERT INTO links (
			short_link,
			alias,
			original_link,
			organization_id,
			created_at,
			expires_in
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
		RETURNING id
	`
	var (
		id string
	)

	err := exec.QueryRow(
		ctx,
		query,
		link.ShortURL(),
		link.Alias(),
		link.OriginalURL(),
		link.OrganizationID(),
		link.CreatedAt(),
		link.ExpiresIn(),
	).Scan(&id)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, domain.ErrShorlinkAlreadyExists
		}

		return nil, fmt.Errorf("postgres create: %w", err)
	}

	return domain.RehydrateLink(domain.RehydrateLinkParams{
		ID:             link.ID(),
		ShortURL:       link.ShortURL(),
		OriginalURL:    link.OriginalURL(),
		OrganizationID: link.OrganizationID(),
		Alias:          link.Alias(),
		ExpiresIn:      link.ExpiresIn(),
		CreatedAt:      link.CreatedAt(),
	})
}

func (r *linkRepository) Update(ctx context.Context, link *domain.Link) (*domain.Link, error) {
	exec := executor(ctx, r.pool)
	query := `
		UPDATE links
		SET
			short_link = $1,
			alias = $2,
			original_link = $3,
			organization_id = $4,
			expires_in = $5,
			updated_at = $6
		WHERE id = $7
	`
	_, err := exec.Exec(
		ctx,
		query,
		link.ShortURL(),
		link.Alias(),
		link.OriginalURL(),
		link.OrganizationID(),
		link.ExpiresIn(),
		time.Now(),
		link.ID(),
	)

	if err != nil {
		return nil, fmt.Errorf("postgres update: %w", err)
	}

	return nil, nil
}

func (r *linkRepository) FindByID(ctx context.Context, linkID string) (*domain.Link, error) {
	exec := executor(ctx, r.pool)
	query := `
		SELECT
			id,
			short_link,
			alias,
			original_link,
			organization_id,
			created_at,
			expires_in,
			updated_at
		FROM links
		WHERE id = $1
	`

	var (
		id              string
		short_link      string
		alias           string
		original_link   string
		organization_id string
		created_at      time.Time
		expires_in      *time.Time
		updated_at      *time.Time
	)

	err := exec.QueryRow(
		ctx,
		query,
		linkID,
	).Scan(
		&id,
		&short_link,
		&alias,
		&original_link,
		&organization_id,
		&created_at,
		&expires_in,
		&updated_at,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrLinkNotFound
		}
		return nil, fmt.Errorf("postgres find by id: %w", err)
	}

	return domain.RehydrateLink(domain.RehydrateLinkParams{
		ID:             id,
		ShortURL:       short_link,
		OriginalURL:    original_link,
		OrganizationID: organization_id,
		Alias:          alias,
		ExpiresIn:      expires_in,
		CreatedAt:      created_at,
		UpdatedAt:      updated_at,
	})
}

func (r *linkRepository) FindByShortLink(ctx context.Context, shortLink string) (*domain.Link, error) {
	exec := executor(ctx, r.pool)
	query := `
		SELECT
			id,
			short_link,
			alias,
			original_link,
			organization_id,
			created_at,
			expires_in,
			updated_at
		FROM links
		WHERE short_link = $1
	`

	var (
		id              string
		short_link      string
		alias           string
		original_link   string
		organization_id string
		created_at      time.Time
		expires_in      *time.Time
		updated_at      *time.Time
	)

	err := exec.QueryRow(
		ctx,
		query,
		shortLink,
	).Scan(
		&id,
		&short_link,
		&alias,
		&original_link,
		&organization_id,
		&created_at,
		&expires_in,
		&updated_at,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrLinkNotFound
		}
		return nil, fmt.Errorf("postgres find by short link: %w", err)
	}

	return domain.RehydrateLink(domain.RehydrateLinkParams{
		ID:             id,
		ShortURL:       short_link,
		OriginalURL:    original_link,
		OrganizationID: organization_id,
		Alias:          alias,
		ExpiresIn:      expires_in,
		CreatedAt:      created_at,
		UpdatedAt:      updated_at,
	})
}
