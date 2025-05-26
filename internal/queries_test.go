package database

import (
	"context"
	"testing"

	"go_magic_probs_tool/internal/models"

	"github.com/pashagolub/pgxmock/v2"
	"github.com/stretchr/testify/assert"
)

func TestFetchBoosterVariants(t *testing.T) {
	// Create a mock database connection
	mockPool, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("Failed to create mock connection: %v", err)
	}
	defer mockPool.Close()

	// Mocking the query result
	rows := pgxmock.NewRows([]string{
		"boostername", "boosterindex", "setcode", "sheetname", "sheetpicks", "boosterweight", "booster_probability",
	}).
		AddRow("BoosterA", new(int), "SET1", "Sheet1", 3, 0.5, 0.25).
		AddRow("BoosterB", new(int), "SET2", "Sheet2", 4, 0.3, 0.15)

	// Set the expected query and rows
	mockPool.ExpectQuery("SELECT c.boostername").
		WithArgs([]string{"SET1", "SET2"}, []string{"BoosterA", "BoosterB"}).
		WillReturnRows(rows)

	// Initialize your DatabaseHandler with the mock pool
	dbHandler := &DatabaseHandler{
		Pool: mockPool,
	}

	// Perform the test
	setCodes := []string{"SET1", "SET2"}
	boosterNames := []string{"BoosterA", "BoosterB"}

	result, err := dbHandler.FetchBoosterVariants(context.Background(), setCodes, boosterNames)
	assert.NoError(t, err)
	assert.Len(t, result, 2)

	// Check the fetched data
	expected := []models.BoosterVariant{
		{
			BoosterName:        "BoosterA",
			BoosterIndex:       new(int),
			SetCode:            "SET1",
			SheetName:          "Sheet1",
			SheetPicks:         3,
			BoosterWeight:      0.5,
			BoosterProbability: 0.25,
		},
		{
			BoosterName:        "BoosterB",
			BoosterIndex:       new(int),
			SetCode:            "SET2",
			SheetName:          "Sheet2",
			SheetPicks:         4,
			BoosterWeight:      0.3,
			BoosterProbability: 0.15,
		},
	}

	assert.Equal(t, expected, result)

	// Ensure all expectations were met
	err = mockPool.ExpectationsWereMet()
	assert.NoError(t, err)
}
