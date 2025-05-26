package cards

import (
	"context"
)

type CardFetcher interface {
	FetchCardData(ctx context.Context, cardNames []string, limit int) (map[string]CardData, error)
}
