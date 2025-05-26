package boosters

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, boosterService *BoosterService) {
	r.HandleFunc("/boosters", GetBoosterVariantsHandler(boosterService)).Methods("POST")
	// add other calculate routes here
}
