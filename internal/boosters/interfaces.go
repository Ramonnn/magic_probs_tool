package boosters

import (
	"context"
)

// BoosterFetcher defines an interface for fetching booster variants
type BoosterFetcher interface {
	FetchBoosterVariants(ctx context.Context, setCodes []string, boosterNames []string) ([]BoosterVariant, error)
}
