package calculate

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type CalculateRequest struct {
	Cards []string `json:"cards"`
}

type CalculateResponse struct {
	Probabilities map[string][]CardProbability `json:"probabilities"`
	CardData      any                          `json:"cardData"`
	Error         string                       `json:"error,omitempty"`
}

func calculateHandler(probService *ProbabilitiesService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		var req CalculateRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		cardProbabilities, cardData, err := probService.GetCardProbabilities(ctx, req.Cards)
		if err != nil {
			log.Printf("calculation error: %v", err)
			json.NewEncoder(w).Encode(CalculateResponse{Error: err.Error()})
			return
		}

		// Flatten to map[cardID][]CardProbability
		flatProbs := make(map[string][]CardProbability)
		for boosterName, cardMap := range cardProbabilities {
			for cardID, probList := range cardMap {
				for _, p := range probList {
					entry := CardProbability{
						Probability: p.Probability,
						IsFoil:      p.IsFoil,
						SetCode:     p.SetCode,
						BoosterName: boosterName, // NEW FIELD
					}
					flatProbs[cardID] = append(flatProbs[cardID], entry)
				}
			}
		}

		json.NewEncoder(w).Encode(CalculateResponse{
			Probabilities: flatProbs,
			CardData:      cardData,
		})
	}
}
