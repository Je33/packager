package packer

import (
	"context"
	"fmt"

	"github.com/Je33/packager/internal/domain"
)

// Delete deletes a package type by UID
func (s *Service) Delete(ctx context.Context, req domain.PackDeleteRequest) (*domain.PackDeleteResponse, error) {
	pack, err := s.repo.PackGetOne(ctx, req.UID)
	if err != nil {
		return nil, fmt.Errorf("getting pack error: %w", err)
	}

	err = s.repo.PackDelete(ctx, req.UID)
	if err != nil {
		return nil, fmt.Errorf("deleting pack error: %w", err)
	}

	return &domain.PackDeleteResponse{
		Pack: pack,
	}, nil
}
