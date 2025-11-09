package packer

import (
	"context"
	"fmt"
	"sort"

	"github.com/Je33/packager/internal/domain"
)

// dpState represents the state for a specific order amount
// Uses backpointers for memory efficiency instead of copying pack maps
type dpState struct {
	minItems   int // minimum items needed
	minPacks   int // minimum packs needed (among solutions with minItems)
	packSize   int // which pack size was used to reach this state
	prevAmount int // previous amount (for backtracking)
}

func (s *Service) Calculate(ctx context.Context, req domain.PackCalculateRequest) (*domain.PackCalculateResponse, error) {
	if req.Items <= 0 {
		return &domain.PackCalculateResponse{
			Calculations: []*domain.Calculation{},
		}, nil
	}

	// Get all available pack sizes from repository
	packs, err := s.repo.PackGetAll(ctx, domain.PackGetAllRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed to get packs: %w", err)
	}

	if len(packs) == 0 {
		return nil, fmt.Errorf("no pack sizes available")
	}

	// Extract pack sizes and sort them
	packSizes := make([]int, len(packs))
	packMap := make(map[int]*domain.Pack) // size -> pack details
	for i, pack := range packs {
		packSizes[i] = pack.Size
		packMap[pack.Size] = pack
	}
	sort.Ints(packSizes)

	// Calculate optimal pack combination
	dp, bestIdx := calculateOptimalPacks(req.Items, packSizes)
	if bestIdx == -1 {
		return nil, fmt.Errorf("unable to fulfill order with available pack sizes")
	}

	// Reconstruct solution from backpointers
	packCounts := reconstructSolution(dp, bestIdx)

	// Convert result to response format
	calculations := make([]*domain.Calculation, 0, len(packCounts))
	for size, quantity := range packCounts {
		if quantity > 0 {
			pack := packMap[size]
			calculations = append(calculations, &domain.Calculation{
				PackUID:  pack.UID,
				PackSize: pack.Size,
				Quantity: quantity,
				Items:    pack.Size * quantity,
			})
		}
	}

	return &domain.PackCalculateResponse{
		Calculations: calculations,
	}, nil
}

// calculateOptimalPacks uses dynamic programming to find optimal pack combination
// Rule 1: Only whole packs
// Rule 2: Minimize total items (primary objective)
// Rule 3: Minimize number of packs (secondary objective, rule 2 takes precedence)
//
// Returns: dp table and best index (-1 if no solution found)
// Uses backpointers instead of copying maps for O(1) memory per state
func calculateOptimalPacks(orderItems int, packSizes []int) ([]*dpState, int) {
	// We need to check up to orderItems + the largest pack size
	// to ensure we can find solutions that may overshoot
	maxPackSize := packSizes[len(packSizes)-1]
	maxAmount := orderItems + maxPackSize

	// dp[i] = best solution to fulfill exactly i items
	dp := make([]*dpState, maxAmount+1)
	dp[0] = &dpState{
		minItems:   0,
		minPacks:   0,
		packSize:   0,
		prevAmount: -1,
	}

	// Track the best solution found so far for early termination
	bestIdx := -1
	bestItems := int(^uint(0) >> 1) // max int

	// Fill DP table using bottom-up dynamic programming
	// For each amount i, try all pack sizes and keep the best solution
	for i := 1; i <= maxAmount; i++ {
		for _, packSize := range packSizes {
			prev := i - packSize
			if prev >= 0 && dp[prev] != nil {
				newItems := dp[prev].minItems + packSize
				newPacks := dp[prev].minPacks + 1

				// Update if this solution is better:
				// 1st priority: fewer total items (Rule #2)
				// 2nd priority: fewer packs (Rule #3)
				if dp[i] == nil ||
					newItems < dp[i].minItems ||
					(newItems == dp[i].minItems && newPacks < dp[i].minPacks) {

					dp[i] = &dpState{
						minItems:   newItems,
						minPacks:   newPacks,
						packSize:   packSize, // store which pack was used
						prevAmount: prev,     // store previous state for backtracking
					}
				}
			}
		}

		// Early termination: track the best solution for items >= orderItems
		if i >= orderItems && dp[i] != nil {
			if dp[i].minItems < bestItems ||
				(dp[i].minItems == bestItems && dp[i].minPacks < dp[bestIdx].minPacks) {
				bestIdx = i
				bestItems = dp[i].minItems
			}

			// Perfect match found - we can stop early
			if i == orderItems && dp[i].minItems == orderItems {
				return dp, bestIdx
			}
		}
	}

	return dp, bestIdx
}

// reconstructSolution rebuilds the pack counts from backpointers
// This avoids O(n*m) map copies during DP, doing only one O(m) reconstruction
func reconstructSolution(dp []*dpState, idx int) map[int]int {
	if idx == -1 || dp[idx] == nil {
		return nil
	}

	// Preallocate with reasonable capacity
	packCounts := make(map[int]int, 8)

	// Backtrack through the solution
	current := idx
	for current > 0 && dp[current] != nil {
		state := dp[current]
		packCounts[state.packSize]++
		current = state.prevAmount
	}

	return packCounts
}
