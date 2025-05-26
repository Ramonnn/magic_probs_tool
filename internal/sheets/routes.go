package sheets

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, sheetService *SheetService) {
	r.HandleFunc("/sheets", GetBoosterSheetsHandler(sheetService)).Methods("POST")
	// add other calculate routes here
}
