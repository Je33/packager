package mem

import (
	"context"

	"github.com/Je33/packager/internal/domain"
)

func (r *Repository) PackGetAll(ctx context.Context, params domain.PackGetAllRequest) ([]*domain.Pack, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	size := len(r.packs)
	if params.Limit > 0 && params.Limit < size {
		size = params.Limit
	}

	packs := make([]*domain.Pack, 0, size)
	for i, pack := range r.packs {
		if params.Page > 0 {
			if i < (params.Page-1)*params.Limit {
				continue
			}
		}

		packs = append(packs, pack)

		if params.Limit > 0 && len(packs) >= params.Limit {
			break
		}
	}

	return packs, nil
}

func (r *Repository) PackGetOne(ctx context.Context, uid string) (*domain.Pack, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, ok := r.packsInd[uid]
	if !ok {
		return nil, domain.ErrPackNotFound
	}

	return r.packs[r.packsInd[uid]], nil
}

func (r *Repository) PackCreate(ctx context.Context, pack *domain.Pack) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.packsInd[pack.UID] = len(r.packs)
	r.packs = append(r.packs, pack)

	return nil
}

func (r *Repository) PackUpdate(ctx context.Context, pack *domain.Pack) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.packsInd[pack.UID]
	if !ok {
		return domain.ErrPackNotFound
	}

	r.packs[r.packsInd[pack.UID]] = &domain.Pack{
		UID:  pack.UID,
		Size: pack.Size,
	}

	return nil
}

func (r *Repository) PackDelete(ctx context.Context, uid string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	idx, ok := r.packsInd[uid]
	if !ok {
		return domain.ErrPackNotFound
	}

	// Remove from slice
	r.packs = append(r.packs[:idx], r.packs[idx+1:]...)

	// Remove from index map
	delete(r.packsInd, uid)

	// Update indices for all packs after the deleted one
	for k, v := range r.packsInd {
		if v > idx {
			r.packsInd[k] = v - 1
		}
	}

	return nil
}
