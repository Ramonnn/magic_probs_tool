package boosters

import (
	"encoding/json"
	"net/http"
	"strings"
)

// GetCardDataHandler handles HTTP requests to fetch card data.
func GetBoosterVariantsHandler(boosterService *BoosterService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract card names from query parameters
		setCodesParam := r.URL.Query().Get("setCodes")
		if setCodesParam == "" {
			http.Error(w, "Missing setCodes query parameter", http.StatusBadRequest)
			return
		}

		setCodes := strings.Split(setCodesParam, ",")

		boosterNamesParam := r.URL.Query().Get("boosterNames")
		if boosterNamesParam == "" {
			http.Error(w, "Missing boosterNames query parameter", http.StatusBadRequest)
			return
		}
		// Split card names by comma
		boosterNames := strings.Split(boosterNamesParam, ",")

		// Fetch card data from the database
		variants, err := boosterService.FetchBoosterVariants(r.Context(), setCodes, boosterNames)
		if err != nil {
			http.Error(w, "Failed to fetch cards", http.StatusInternalServerError)
			return
		}

		// Set response header and write JSON response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(variants)
	}
}
