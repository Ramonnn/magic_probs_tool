package calculate

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func calculateHandler(probService *ProbabilitiesService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		var req CalculateRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		cardProbabilities, err := probService.GetCardProbabilities(ctx, req.Cards)
		if err != nil {
			log.Printf("calculation error: %v", err)
			json.NewEncoder(w).Encode(CalculateResponse{Error: err.Error()})
			return
		}

		// Flatten to map[cardID][]CardProbability
		flatProbs := make(map[string][]CardProbability)
		for _, cardMap := range cardProbabilities {
			for cardID, probList := range cardMap {
				flatProbs[cardID] = append(flatProbs[cardID], probList...)
			}
		}

		// Create a slice of Row for aggregation
		var allRows []Row
		for cardID, probs := range flatProbs {
			for _, p := range probs {
				allRows = append(allRows, Row{
					UUID:           cardID,
					Booster:        p.BoosterName,
					BoosterVariant: p.BoosterVariant,
					Set:            p.SetCode,
					Foil:           p.IsFoil,
					PromoTypes:     p.PromoTypes,
					FrameEffects:   p.FrameEffects,
					Sheet:          p.SheetName,
					SheetPicks:     p.SheetPicks,
					Probability:    p.Probability,
				})
			}
		}

		// Aggregate by booster
		aggByBooster := AggregateBy(allRows, func(r Row) string {
			return r.Booster
		})

		// Aggregate by booster+foil
		aggByBoosterFoil := AggregateBy(allRows, func(r Row) string {
			return r.Booster + "|" + fmt.Sprint(r.Foil)
		})

		// Aggregate by set+booster
		aggBySetBooster := AggregateBy(allRows, func(r Row) string {
			return r.Set + "|" + r.Booster
		})

		// Intermediate nested map: set → booster → []rows
		grouped := make(map[string]map[string][]Row)
		for _, row := range allRows {
			if _, ok := grouped[row.Set]; !ok {
				grouped[row.Set] = make(map[string][]Row)
			}
			grouped[row.Set][row.Booster] = append(grouped[row.Set][row.Booster], row)
		}

		// Convert to response structure
		var drilldown []DrilldownSet
		for setCode, boosterMap := range grouped {
			var setTotal float64
			var boosters []DrilldownBooster

			for boosterName, rows := range boosterMap {
				var boosterTotal float64
				var cards []DrilldownCard

				for _, row := range rows {
					boosterTotal += row.Probability
					cards = append(cards, DrilldownCard{
						UUID:         row.UUID,
						Foil:         row.Foil,
						PromoTypes:   row.PromoTypes,
						FrameEffects: row.FrameEffects,
						Probability:  row.Probability,
					})
				}

				setTotal += boosterTotal

				boosters = append(boosters, DrilldownBooster{
					BoosterName: boosterName,
					Cards:       cards,
					TotalProb:   boosterTotal,
				})
			}

			drilldown = append(drilldown, DrilldownSet{
				SetCode:   setCode,
				Boosters:  boosters,
				TotalProb: setTotal,
			})
		}

		json.NewEncoder(w).Encode(CalculateResponse{
			Probabilities:       flatProbs,
			AggregatedByBooster: aggByBooster,
			AggregatedByFoil:    aggByBoosterFoil,
			AggregatedBySet:     aggBySetBooster,
			Drilldown:           drilldown,
		})
	}
}
