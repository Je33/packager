package resolver

import (
	"context"

	"github.com/Je33/packager/internal/domain"
)

//go:generate go tool gqlgen generate

type Packer interface {
	Calculate(ctx context.Context, req domain.PackCalculateRequest) (*domain.PackCalculateResponse, error)
	Create(ctx context.Context, req domain.PackCreateRequest) (*domain.PackCreateResponse, error)
	Update(ctx context.Context, req domain.PackUpdateRequest) (*domain.PackUpdateResponse, error)
	Delete(ctx context.Context, req domain.PackDeleteRequest) (*domain.PackDeleteResponse, error)
	GetAll(ctx context.Context, req domain.PackGetAllRequest) (*domain.PackGetAllResponse, error)
}

type Resolver struct {
	packer Packer
}

func New(packer Packer) *Resolver {
	return &Resolver{
		packer: packer,
	}
}
