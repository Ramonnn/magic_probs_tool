package calculate

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, probService *ProbabilitiesService) {
	r.HandleFunc("/api/calculate", calculateHandler(probService)).Methods("POST")
}
