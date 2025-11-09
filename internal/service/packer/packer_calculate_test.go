package packer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Je33/packager/internal/domain"
)

func TestCalculate(t *testing.T) {
	// Standard pack sizes from requirements
	standardPacks := []*domain.Pack{
		{UID: "250", Size: 250},
		{UID: "500", Size: 500},
		{UID: "1000", Size: 1000},
		{UID: "2000", Size: 2000},
		{UID: "5000", Size: 5000},
	}

	tests := []struct {
		name          string
		items         int
		packs         []*domain.Pack
		wantTotal     int
		wantPackCount int
		wantPacks     map[int]int // size -> quantity
	}{
		{
			name:          "1 item",
			items:         1,
			packs:         standardPacks,
			wantTotal:     250,
			wantPackCount: 1,
			wantPacks:     map[int]int{250: 1},
		},
		{
			name:          "250 items",
			items:         250,
			packs:         standardPacks,
			wantTotal:     250,
			wantPackCount: 1,
			wantPacks:     map[int]int{250: 1},
		},
		{
			name:          "251 items - should use 1x500 not 2x250",
			items:         251,
			packs:         standardPacks,
			wantTotal:     500,
			wantPackCount: 1,
			wantPacks:     map[int]int{500: 1},
		},
		{
			name:          "501 items - should use 1x500 + 1x250",
			items:         501,
			packs:         standardPacks,
			wantTotal:     750,
			wantPackCount: 2,
			wantPacks:     map[int]int{500: 1, 250: 1},
		},
		{
			name:          "12001 items - should use 2x5000 + 1x2000 + 1x250",
			items:         12001,
			packs:         standardPacks,
			wantTotal:     12250,
			wantPackCount: 4,
			wantPacks:     map[int]int{5000: 2, 2000: 1, 250: 1},
		},
		{
			name:  "Edge case with 23, 31, 53 pack sizes",
			items: 500000,
			packs: []*domain.Pack{
				{UID: "23", Size: 23},
				{UID: "31", Size: 31},
				{UID: "53", Size: 53},
			},
			wantTotal:     500000, // 23*2 + 31*7 + 53*9429 = 46 + 217 + 499737 = 500000 exactly!
			wantPackCount: 9438,   // 2 + 7 + 9429 = 9438
			wantPacks:     map[int]int{23: 2, 31: 7, 53: 9429},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := NewMockRepository(ctrl)
			mockRepo.EXPECT().PackGetAll(gomock.Any(), gomock.Any()).Return(tt.packs, nil)

			service := &Service{
				repo: mockRepo,
			}

			resp, err := service.Calculate(context.Background(), domain.PackCalculateRequest{
				Items: tt.items,
			})

			assert.NoError(t, err)
			assert.NotNil(t, resp)

			// Calculate total items and pack count
			totalItems := 0
			totalPacks := 0
			actualPacks := make(map[int]int)

			for _, calc := range resp.Calculations {
				totalItems += calc.PackSize * calc.Quantity
				totalPacks += calc.Quantity
				actualPacks[calc.PackSize] = calc.Quantity
			}

			// Verify total items matches expected
			assert.Equal(t, tt.wantTotal, totalItems, "Total items should match")

			// Verify pack count matches expected
			assert.Equal(t, tt.wantPackCount, totalPacks, "Total pack count should match")

			// Verify exact pack distribution
			assert.Equal(t, tt.wantPacks, actualPacks, "Pack distribution should match")

			// Verify we're sending at least the required items
			assert.GreaterOrEqual(t, totalItems, tt.items, "Should send at least requested items")
		})
	}
}

func TestCalculate_EmptyRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	service := &Service{repo: mockRepo}

	resp, err := service.Calculate(context.Background(), domain.PackCalculateRequest{
		Items: 0,
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Empty(t, resp.Calculations)
}

func TestCalculate_NoPacks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockRepo.EXPECT().PackGetAll(gomock.Any(), gomock.Any()).Return([]*domain.Pack{}, nil)

	service := &Service{repo: mockRepo}

	resp, err := service.Calculate(context.Background(), domain.PackCalculateRequest{
		Items: 100,
	})

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no pack sizes available")
}

func TestCalculate_AdditionalEdgeCases(t *testing.T) {
	standardPacks := []*domain.Pack{
		{UID: "250", Size: 250},
		{UID: "500", Size: 500},
		{UID: "1000", Size: 1000},
		{UID: "2000", Size: 2000},
		{UID: "5000", Size: 5000},
	}

	tests := []struct {
		name  string
		items int
		// We'll verify the solution satisfies the rules rather than exact packs
		checkFunc func(t *testing.T, resp *domain.PackCalculateResponse, items int)
	}{
		{
			name:  "749 items - just under 750",
			items: 749,
			checkFunc: func(t *testing.T, resp *domain.PackCalculateResponse, items int) {
				total := 0
				for _, c := range resp.Calculations {
					total += c.PackSize * c.Quantity
				}
				assert.GreaterOrEqual(t, total, items)
				// Should be 750 (500+250) not 1000
				assert.LessOrEqual(t, total, 750)
			},
		},
		{
			name:  "Large order - 50000",
			items: 50000,
			checkFunc: func(t *testing.T, resp *domain.PackCalculateResponse, items int) {
				total := 0
				packCount := 0
				for _, c := range resp.Calculations {
					total += c.PackSize * c.Quantity
					packCount += c.Quantity
				}
				assert.GreaterOrEqual(t, total, items)
				// 50000 exactly = 10 x 5000
				assert.Equal(t, 50000, total)
				assert.Equal(t, 10, packCount)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := NewMockRepository(ctrl)
			mockRepo.EXPECT().PackGetAll(gomock.Any(), gomock.Any()).Return(standardPacks, nil)

			service := &Service{repo: mockRepo}
			resp, err := service.Calculate(context.Background(), domain.PackCalculateRequest{
				Items: tt.items,
			})

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			tt.checkFunc(t, resp, tt.items)
		})
	}
}

// Test the core algorithm directly
func TestCalculateOptimalPacks(t *testing.T) {
	tests := []struct {
		name          string
		orderItems    int
		packSizes     []int
		wantTotal     int
		wantPackCount int
		wantPacks     map[int]int
	}{
		{
			name:          "Simple case - exact match",
			orderItems:    250,
			packSizes:     []int{250, 500, 1000, 2000, 5000},
			wantTotal:     250,
			wantPackCount: 1,
			wantPacks:     map[int]int{250: 1},
		},
		{
			name:          "251 items - minimize items first",
			orderItems:    251,
			packSizes:     []int{250, 500, 1000, 2000, 5000},
			wantTotal:     500,
			wantPackCount: 1,
			wantPacks:     map[int]int{500: 1},
		},
		{
			name:          "501 items - minimize items then packs",
			orderItems:    501,
			packSizes:     []int{250, 500, 1000, 2000, 5000},
			wantTotal:     750,
			wantPackCount: 2,
			wantPacks:     map[int]int{500: 1, 250: 1},
		},
		{
			name:          "Edge case: 500000 with [23, 31, 53]",
			orderItems:    500000,
			packSizes:     []int{23, 31, 53},
			wantTotal:     500000,
			wantPackCount: 9438,
			wantPacks:     map[int]int{23: 2, 31: 7, 53: 9429},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dp, bestIdx := calculateOptimalPacks(tt.orderItems, tt.packSizes)

			assert.NotEqual(t, -1, bestIdx, "Should find a solution")
			assert.NotNil(t, dp[bestIdx])

			result := dp[bestIdx]
			packCounts := reconstructSolution(dp, bestIdx)

			assert.Equal(t, tt.wantTotal, result.minItems, "Total items should match")
			assert.Equal(t, tt.wantPackCount, result.minPacks, "Pack count should match")
			assert.Equal(t, tt.wantPacks, packCounts, "Pack distribution should match")
		})
	}
}

// Benchmark the edge case to ensure performance
func BenchmarkCalculateOptimalPacks_EdgeCase(b *testing.B) {
	packSizes := []int{23, 31, 53}
	orderItems := 500000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dp, bestIdx := calculateOptimalPacks(orderItems, packSizes)
		if bestIdx == -1 {
			b.Fatal("Expected result, got no solution")
		}
		_ = reconstructSolution(dp, bestIdx)
	}
}

// Benchmark standard case
func BenchmarkCalculateOptimalPacks_Standard(b *testing.B) {
	packSizes := []int{250, 500, 1000, 2000, 5000}
	orderItems := 12001

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dp, bestIdx := calculateOptimalPacks(orderItems, packSizes)
		if bestIdx == -1 {
			b.Fatal("Expected result, got no solution")
		}
		_ = reconstructSolution(dp, bestIdx)
	}
}
