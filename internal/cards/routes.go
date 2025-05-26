package cards

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, cardService *CardService) {
	r.HandleFunc("/cards", GetCardDataHandler(cardService)).Methods("POST")
	// add other calculate routes here
}
