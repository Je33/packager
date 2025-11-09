package packer

import (
	"context"
	"fmt"

	"github.com/Je33/packager/internal/domain"
)

// GetAll returns all package types
func (s *Service) GetAll(ctx context.Context, req domain.PackGetAllRequest) (*domain.PackGetAllResponse, error) {
	packs, err := s.repo.PackGetAll(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("getting packs error: %w", err)
	}

	return &domain.PackGetAllResponse{
		Packs: packs,
	}, nil
}
