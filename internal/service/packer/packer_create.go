package packer

import (
	"context"
	"fmt"

	"github.com/Je33/packager/internal/domain"
	"github.com/Je33/packager/pkg/uid"
)

// Create creates a new package type
func (s *Service) Create(ctx context.Context, req domain.PackCreateRequest) (*domain.PackCreateResponse, error) {
	pack := &domain.Pack{
		UID:  uid.Gen(8),
		Size: req.Size,
	}

	err := s.repo.PackCreate(ctx, pack)
	if err != nil {
		return nil, fmt.Errorf("creating pack error: %w", err)
	}

	return &domain.PackCreateResponse{
		Pack: pack,
	}, nil
}
