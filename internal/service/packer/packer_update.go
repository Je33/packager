package packer

import (
	"context"
	"fmt"

	"github.com/Je33/packager/internal/domain"
)

// Update updates a package type by UID
func (s *Service) Update(ctx context.Context, req domain.PackUpdateRequest) (*domain.PackUpdateResponse, error) {
	pack, err := s.repo.PackGetOne(ctx, req.UID)
	if err != nil {
		return nil, fmt.Errorf("getting pack error: %w", err)
	}

	pack.Size = req.Size

	err = s.repo.PackUpdate(ctx, pack)
	if err != nil {
		return nil, fmt.Errorf("updating pack error: %w", err)
	}

	return &domain.PackUpdateResponse{
		Pack: pack,
	}, nil
}
