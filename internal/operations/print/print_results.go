package print

import (
	"fmt"
	"go_magic_probs_tool/internal/cards"
	"go_magic_probs_tool/internal/operations/calculate"
	"sort"
	"strings"
)

const separatorPadding = 4

type CardEntry struct {
	UUID string
	Data calculate.CardProbability
}

type CardEntries []CardEntry

func (ce CardEntries) Len() int      { return len(ce) }
func (ce CardEntries) Swap(i, j int) { ce[i], ce[j] = ce[j], ce[i] }
func (ce CardEntries) Less(i, j int) bool {
	return ce[i].Data.Probability > ce[j].Data.Probability
}

func PrintCardProbabilities(cardProbabilities map[string]map[string][]calculate.CardProbability, cardData map[string]cards.CardData) {
	// Define column widths
	columnWidths := map[string]int{
		"name":        60,
		"set_code":    10,
		"foil":        8,
		"effects":     20,
		"promo":       12,
		"probability": 20,
	}

	// Print header row
	fmt.Printf("%-*s %-*s %-*s %-*s %-*s %-*s\n",
		columnWidths["name"], "Card Name",
		columnWidths["set_code"], "Set Code",
		columnWidths["foil"], "Foil",
		columnWidths["effects"], "Frame Effects",
		columnWidths["promo"], "Promo Types",
		columnWidths["probability"], "Probability",
	)
	fmt.Println(strings.Repeat("=", sum(columnWidths)+separatorPadding))

	// Iterate over each booster
	for boosterName, cardProbs := range cardProbabilities {
		if len(cardProbs) == 0 {
			// No cards to print for this booster, skip it
			continue
		}

		fmt.Printf("\n🎲 Card Probabilities in %s Booster:\n", boosterName)

		// Flatten and sort cards by probability
		sortedCards := make(CardEntries, 0)
		for cardUUID, variants := range cardProbs {
			for _, variant := range variants {
				sortedCards = append(sortedCards, CardEntry{UUID: cardUUID, Data: variant})
			}
		}

		sort.Sort(sortedCards)

		for _, cardEntry := range sortedCards {
			cardUUID := cardEntry.UUID
			data := cardEntry.Data
			cardInfo := cardData[cardUUID]

			// Extract data
			cardName := cardInfo.Name
			setCode := data.SetCode
			foilStatus := "No Foil"
			if data.IsFoil {
				foilStatus = "(Foil)"
			}

			frameEffects := "N/A"
			if cardInfo.FrameEffects != nil {
				frameEffects = strings.Join(*cardInfo.FrameEffects, ", ")
			}
			promoTypes := "N/A"
			if cardInfo.PromoTypes != nil {
				promoTypes = strings.Join(*cardInfo.PromoTypes, ", ")
			}

			// Print each row
			fmt.Printf("%-*s %-*s %-*s %-*s %-*s %-*s\n",
				columnWidths["name"], cardName,
				columnWidths["set_code"], setCode,
				columnWidths["foil"], foilStatus,
				columnWidths["effects"], frameEffects,
				columnWidths["promo"], promoTypes,
				columnWidths["probability"], fmt.Sprintf("%.4f%%", data.Probability*100),
			)
		}
	}
}

func sum(m map[string]int) int {
	total := 0
	for _, v := range m {
		total += v
	}
	return total
}
