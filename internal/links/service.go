package links

import (
	"context"
	"fmt"

	"github.com/cleidison-barradas/shortr.api/internal/app"
	"github.com/cleidison-barradas/shortr.api/internal/domain"
	"github.com/cleidison-barradas/shortr.api/internal/utils"
)

type linkService struct {
	linkRepo app.LinkRepository
}

func newLinksService(
	linkRepo app.LinkRepository,
) *linkService {
	return &linkService{
		linkRepo: linkRepo,
	}
}

func (s *linkService) FindByShortURL(ctx context.Context, shortLink string) (*LinkResponse, error) {
	result, err := s.linkRepo.FindByShortLink(ctx, shortLink)
	if err != nil {
		return nil, err
	}

	isExpired := result.IsExpired()
	if isExpired {
		return nil, app.ErrLinkExpired
	}

	return &LinkResponse{
		ID:          result.ID(),
		ShortURL:    result.ShortURL(),
		OriginalURL: result.OriginalURL(),
		Alias:       result.Alias(),
		ExpiresIn:   result.ExpiresIn(),
		CreatedAt:   result.CreatedAt(),
		UpdatedAt:   result.UpdatedAt(),
	}, nil
}

func (s *linkService) CreateLink(ctx context.Context, req LinkRequest) (*LinkResponse, error) {
	shortURL, err := utils.GenerateShortCode(4)
	if err != nil {
		return nil, fmt.Errorf("GenerateShortCode error: %w", err)
	}

	data, err := domain.NewLink(domain.NewLinkParams{
		ShortURL:       shortURL,
		Alias:          req.Alias,
		OriginalURL:    req.OriginalURL,
		ExpiresIn:      req.ExpiresIn,
		OrganizationID: req.OrgID,
	})

	link, err := s.linkRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return &LinkResponse{
		ID:          link.ID(),
		ShortURL:    link.ShortURL(),
		OriginalURL: link.OriginalURL(),
		Alias:       link.Alias(),
		ExpiresIn:   link.ExpiresIn(),
		CreatedAt:   link.CreatedAt(),
		UpdatedAt:   link.UpdatedAt(),
	}, nil
}
