package mem

import (
	"sync"

	"github.com/Je33/packager/internal/domain"
	"github.com/Je33/packager/pkg/logger"
)

// Repository is an in-memory implementation of Repository
// TODO: implement persistence
type Repository struct {
	mu       sync.RWMutex
	packs    []*domain.Pack
	packsInd map[string]int
	log      logger.Logger
}

func New(log logger.Logger) *Repository {
	return &Repository{
		log: log,
		packs: []*domain.Pack{
			{
				UID:  "250",
				Size: 250,
			},
			{
				UID:  "500",
				Size: 500,
			},
			{
				UID:  "1000",
				Size: 1000,
			},
			{
				UID:  "2000",
				Size: 2000,
			},
			{
				UID:  "5000",
				Size: 5000,
			},
		},
		packsInd: map[string]int{
			"250":  0,
			"500":  1,
			"1000": 2,
			"2000": 3,
			"5000": 4,
		},
	}
}
