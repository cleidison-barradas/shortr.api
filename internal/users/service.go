package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/cleidison-barradas/shortr.api/internal/app"
	"github.com/cleidison-barradas/shortr.api/internal/domain"
)

type UserService struct {
	repo   app.UserRepository
	hasher domain.PasswordHasher
}

func NewUserService(
	repo app.UserRepository,
	hasher domain.PasswordHasher,
) *UserService {
	return &UserService{
		repo:   repo,
		hasher: hasher,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req CreateUserRequest) (*CreatedUserResult, error) {
	_, err := s.repo.FindByEmail(ctx, req.Email)

	if err != nil {
		if !errors.Is(err, domain.ErrUserNotFound) {
			return nil, fmt.Errorf("FindUserByEmail error: %w", err)
		}

		email, err := domain.NewEmail(req.Email)
		if err != nil {
			return nil, err
		}

		passwordHash, err := s.hasher.HashPassword(req.Password)
		if err != nil {
			return nil, fmt.Errorf("HashPassword error: %w", err)
		}

		data, err := domain.NewUser(domain.NewUserParams{
			FirstName:    req.FirstName,
			LastName:     req.LastName,
			Email:        email,
			PasswordHash: passwordHash,
		})

		user, err := s.repo.Create(ctx, data)
		if err != nil {
			return nil, err
		}

		return &CreatedUserResult{
			ID:        user.ID(),
			Name:      user.FullName(),
			Email:     user.Email().String(),
			CreatedAt: user.CreatedAt(),
		}, nil
	}

	return nil, domain.ErrUserEmailAlreadyExists
}

func (s *UserService) FindUsers(ctx context.Context) ([]domain.User, error) {
	return s.repo.Find(ctx)
}
