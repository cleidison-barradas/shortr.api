package organizations

import (
	"context"
	"errors"

	"github.com/cleidison-barradas/shortr.api/internal/app"
	"github.com/cleidison-barradas/shortr.api/internal/domain"
)

type organizationService struct {
	userRepo app.UserRepository
	orgRepo  app.OrganizationRepository
}

func NewOrganizationService(
	userRepo app.UserRepository,
	orgRepo app.OrganizationRepository,
) *organizationService {
	return &organizationService{
		userRepo: userRepo,
		orgRepo:  orgRepo,
	}
}

func (s *organizationService) CreateOrganization(ctx context.Context, req OrganizationRequest) (*OrganizationResponse, error) {

	_, err := s.userRepo.FindByID(ctx, req.UserID)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	orgData, err := domain.NewOrganization(&domain.NewOrganizationParams{
		Name:   req.Name,
		UserID: req.UserID,
	})
	if err != nil {
		return nil, err
	}

	org, err := s.orgRepo.Create(ctx, orgData)
	if err != nil {
		return nil, err
	}

	return &OrganizationResponse{
		ID:        org.ID(),
		Name:      org.Name(),
		Enabled:   org.Enabled(),
		CreatedAt: org.CreatedAt(),
	}, nil
}
