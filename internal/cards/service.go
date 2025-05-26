package cards

import (
	"context"
)

type CardFetcher interface {
	FetchCardData(ctx context.Context, cardNames []string, limit int) (map[string][]CardData, error)
}

// CardService handles booster-related operations
type CardService struct {
	Fetcher CardFetcher
}

// NewCardService creates a new BoosterService with a fetcher
func NewCardService(fetcher CardFetcher) *CardService {
	return &CardService{
		Fetcher: fetcher,
	}
}

func (s *CardService) FetchCardData(ctx context.Context, cardNames []string, batchSize int) (map[string][]CardData, error) {
	return s.Fetcher.FetchCardData(ctx, cardNames, batchSize)
}
