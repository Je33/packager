package packer

import (
	"context"

	"github.com/Je33/packager/internal/domain"
	"github.com/Je33/packager/pkg/logger"
)

//go:generate go tool mockgen -source=$GOFILE -destination=packer_repository_mock_test.go -package=packer
type Repository interface {
	PackGetAll(ctx context.Context, params domain.PackGetAllRequest) ([]*domain.Pack, error)
	PackGetOne(ctx context.Context, uid string) (*domain.Pack, error)
	PackCreate(ctx context.Context, pack *domain.Pack) error
	PackUpdate(ctx context.Context, pack *domain.Pack) error
	PackDelete(ctx context.Context, uid string) error
}

type Service struct {
	repo Repository
	log  logger.Logger
}

func New(repo Repository, log logger.Logger) *Service {
	return &Service{
		repo: repo,
		log:  log,
	}
}
