package main

import (
	"log"
	"net/http"
	"os"

	"go_magic_probs_tool/internal/api"
	"go_magic_probs_tool/internal/boosters"
	"go_magic_probs_tool/internal/cards"
	"go_magic_probs_tool/internal/database"
	"go_magic_probs_tool/internal/operations/calculate"
	"go_magic_probs_tool/internal/sheets"
)

func main() {
	dbHandler, err := database.NewDatabaseHandler()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer dbHandler.Close()

	cardRepo := cards.NewCardRepository(dbHandler)
	boosterRepo := boosters.NewBoosterRepository(dbHandler)
	sheetRepo := sheets.NewSheetRepository(dbHandler)

	cardService := cards.NewCardService(cardRepo)
	boosterService := boosters.NewBoosterService(boosterRepo)
	sheetService := sheets.NewSheetService(sheetRepo)

	probService := calculate.NewProbabilitiesService(cardService, boosterService, sheetService)

	router := api.NewRouter(cardService, probService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
