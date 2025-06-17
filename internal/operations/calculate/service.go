package calculate

import (
	"context"
	"go_magic_probs_tool/internal/boosters"
	"go_magic_probs_tool/internal/cards"
	"go_magic_probs_tool/internal/sheets"
)

// BoosterService handles booster-related operations
type ProbabilitiesService struct {
	cardFetcher    cards.CardFetcher
	boosterFetcher boosters.BoosterFetcher
	sheetFetcher   sheets.SheetFetcher
}

// NewBoosterService creates a new BoosterService with a fetcher
func NewProbabilitiesService(cardFetcher cards.CardFetcher, boosterFetcher boosters.BoosterFetcher, sheetFetcher sheets.SheetFetcher) *ProbabilitiesService {
	return &ProbabilitiesService{cardFetcher: cardFetcher, boosterFetcher: boosterFetcher, sheetFetcher: sheetFetcher}
}

func (s *ProbabilitiesService) GetCardProbabilities(ctx context.Context, cardNames []string) (map[string]map[string][]CardProbability, error) {
	cardData, err := s.cardFetcher.FetchCardData(ctx, cardNames, 500)
	if err != nil {
		return nil, err
	}

	var cardUUIDs []string
	for uuid := range cardData {
		cardUUIDs = append(cardUUIDs, uuid)
	}

	boosterSheets, err := s.sheetFetcher.FetchBoosterSheets(ctx, nil, nil, cardUUIDs)
	if err != nil {
		return nil, err
	}

	setCodes := sheets.ExtractSetCodesFromSheets(boosterSheets)

	boosters, err := s.boosterFetcher.FetchBoosterVariants(ctx, setCodes, nil)
	if err != nil {
		return nil, err
	}

	cardProbabilities, err := CalculateCardProbabilities(boosterSheets, boosters, cardData)
	if err != nil {
		return nil, err
	}

	return cardProbabilities, nil
}
