package boosters

import (
	"context"
)

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
