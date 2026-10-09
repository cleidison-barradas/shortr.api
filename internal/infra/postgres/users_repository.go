package postgres

import (
	"context"
	"errors"

	"github.com/cleidison-barradas/shortr.api/internal/app"
	"github.com/cleidison-barradas/shortr.api/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) app.UserRepository {
	return &userRepository{
		pool: pool,
	}
}

func (r *userRepository) Find(ctx context.Context) ([]domain.User, error) {
	return nil, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	exec := executor(ctx, r.pool)

	query := `
		SELECT 
		id, 
		first_name, 
		last_name,
		email, 
		password_hash 
			FROM users 
			WHERE email = $1`

	var (
		id, firstName, lastName, passwordHash string
	)

	err := exec.QueryRow(ctx, query, email).Scan(&id, &firstName, &lastName, &email, &passwordHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	validatedEmail, err := domain.NewEmail(email)
	if err != nil {
		return nil, domain.ErrInvalidEmail
	}

	return domain.RehydrateUser(domain.RehydratedUserParams{
		ID:           id,
		FirstName:    firstName,
		LastName:     lastName,
		Email:        validatedEmail,
		PasswordHash: passwordHash,
	})
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	exec := executor(ctx, r.pool)

	query := `
	INSERT INTO	users (
		first_name, last_name, email, password_hash
	) VALUES ($1, $2, $3, $4) 
	 RETURNING id
	`

	var (
		id string
	)

	err := exec.QueryRow(
		ctx,
		query,
		user.FirstName(),
		user.LastName(),
		user.Email().String(),
		user.PasswordHash(),
	).Scan(&id)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, domain.ErrUserEmailAlreadyExists
		}

		return nil, err
	}

	return domain.RehydrateUser(domain.RehydratedUserParams{
		ID:        id,
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		Email:     user.Email(),
		CreatedAt: user.CreatedAt(),
	})
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	return nil, nil
}

func (r *userRepository) FindByID(ctx context.Context, userID string) (*domain.User, error) {
	exec := executor(ctx, r.pool)

	query := `
		SELECT 
		id, 
		first_name, 
		last_name
			FROM users 
			WHERE id = $1`

	var (
		id, firstName, lastName string
	)

	err := exec.QueryRow(ctx, query, userID).Scan(&id, &firstName, &lastName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return domain.RehydrateUser(domain.RehydratedUserParams{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
	})
}
