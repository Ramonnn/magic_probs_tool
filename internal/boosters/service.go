package boosters

import (
	"context"
)

// BoosterFetcher defines an interface for fetching booster variants
type BoosterFetcher interface {
	FetchBoosterVariants(ctx context.Context, setCodes []string, boosterNames []string) ([]BoosterVariant, error)
}

// BoosterService handles booster-related operations
type BoosterService struct {
	Fetcher BoosterFetcher
}

// NewBoosterService creates a new BoosterService with a fetcher
func NewBoosterService(fetcher BoosterFetcher) *BoosterService {
	return &BoosterService{
		Fetcher: fetcher,
	}
}

// GetBoosterVariants retrieves booster variants using the fetcher
func (s *BoosterService) FetchBoosterVariants(ctx context.Context, setCodes []string, boosterNames []string) ([]BoosterVariant, error) {
	return s.Fetcher.FetchBoosterVariants(ctx, setCodes, boosterNames)
}
