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
		for _, probs := range flatProbs {
			for _, p := range probs {
				allRows = append(allRows, Row{
					Booster:     p.BoosterName,
					Foil:        p.IsFoil,
					Probability: p.Probability,
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

		json.NewEncoder(w).Encode(CalculateResponse{
			Probabilities:       flatProbs,
			AggregatedByBooster: aggByBooster,
			AggregatedByFoil:    aggByBoosterFoil,
		})
	}
}
