package cards

import (
	"encoding/json"
	"net/http"
	"strings"
)

// GetCardDataHandler handles HTTP requests to fetch card data.
func GetCardDataHandler(cardService *CardService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract card names from query parameters
		cardNamesParam := r.URL.Query().Get("cardNames")
		if cardNamesParam == "" {
			http.Error(w, "Missing cardNames query parameter", http.StatusBadRequest)
			return
		}

		// Split card names by comma
		cardNames := strings.Split(cardNamesParam, ",")

		// Fetch card data from the database
		cards, err := cardService.FetchCardData(r.Context(), cardNames, 10)
		if err != nil {
			http.Error(w, "Failed to fetch cards", http.StatusInternalServerError)
			return
		}

		// Set response header and write JSON response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cards)
	}
}
