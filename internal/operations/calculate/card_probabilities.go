package calculate

import (
	"errors"
	"go_magic_probs_tool/internal/boosters"
	"go_magic_probs_tool/internal/cards"
	"go_magic_probs_tool/internal/sheets"
)

func CalculateCardProbabilities(
	sheets map[string]map[string][]sheets.BoosterSheetEntry,
	boosters []boosters.BoosterVariant,
	cardData map[string][]cards.CardData,
) (map[string]map[string][]CardProbability, error) {
	cardProbabilities := make(map[string]map[string][]CardProbability)

	for _, booster := range boosters {
		boosterName := booster.BoosterName
		sheetName := booster.SheetName
		sheetPicks := booster.SheetPicks
		boosterProbability := booster.BoosterProbability
		setCode := booster.SetCode
		boosterIndex := booster.BoosterIndex

		// Defensive check for sheet existence
		if _, exists := sheets[sheetName]; !exists {
			continue
		}

		if _, exists := cardProbabilities[boosterName]; !exists {
			cardProbabilities[boosterName] = make(map[string][]CardProbability)
		}

		// Calculate probabilities for each card in the sheet
		for cardID, entries := range sheets[sheetName] {
			for _, entry := range entries {
				if entry.BoosterName != boosterName {
					continue
				}

				cardProbability := entry.CardProbability * float64(sheetPicks) * boosterProbability

				// Default empty slices
				var promoTypes []string
				var frameEffects []string

				if variants, ok := cardData[cardID]; ok && len(variants) > 0 {
					promoTypes = variants[0].PromoTypes
					frameEffects = variants[0].FrameEffects
				}

				cardProbabilities[boosterName][cardID] = append(
					cardProbabilities[boosterName][cardID],
					CardProbability{
						Probability:    cardProbability,
						IsFoil:         entry.IsFoil,
						SetCode:        setCode,
						PromoTypes:     promoTypes,
						FrameEffects:   frameEffects,
						BoosterName:    boosterName,
						SheetName:      sheetName,
						SheetPicks:     sheetPicks,
						BoosterVariant: boosterIndex,
					},
				)
			}
		}
	}

	// Return an error if no valid combinations were found
	if len(cardProbabilities) == 0 {
		return nil, errors.New("no valid booster-sheet-set_code combinations found")
	}

	return cardProbabilities, nil
}
