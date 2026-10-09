package signin

import (
	"context"

	"github.com/cleidison-barradas/shortr.api/internal/app"
	"github.com/cleidison-barradas/shortr.api/internal/domain"
)

type siginService struct {
	userRepo app.UserRepository
	hasher   domain.PasswordHasher
}

func NewSiginService(
	userRepo app.UserRepository,
	hasher domain.PasswordHasher,
) *siginService {
	return &siginService{
		userRepo: userRepo,
		hasher:   hasher,
	}
}

func (s *siginService) SignInUser(ctx context.Context, req SigninRequest) (*SigninResponse, error) {

	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, app.ErrInvalidPassword
	}

	matchPassword := user.CheckPassword(req.Password, s.hasher)
	if !matchPassword {
		return nil, app.ErrInvalidPassword
	}

	return &SigninResponse{
		UserID: user.ID(),
		Email:  user.Email().String(),
	}, nil
}
