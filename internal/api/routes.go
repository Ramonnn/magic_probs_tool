package api

import (
	"go_magic_probs_tool/internal/cards"
	"go_magic_probs_tool/internal/operations/calculate"
	// "go_magic_probs_tool/internal/sheets"
	// "go_magic_probs_tool/internal/boosters"
	"github.com/gorilla/mux"
)

func NewRouter(
	cardService *cards.CardService,
	calculateService *calculate.ProbabilitiesService,
	// sheetService *sheets.SheetService,
	// boosterService *boosters.BoosterService,
) *mux.Router {
	r := mux.NewRouter()

	cards.RegisterRoutes(r, cardService)
	calculate.RegisterRoutes(r, calculateService)
	// TODO: figure out if I ever need specific sheet api or booster api functionality.
	// sheets.RegisterRoutes(r, sheetService)
	// boosters.RegisterRoutes(r, boosterService)

	return r
}
