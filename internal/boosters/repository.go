package boosters

import (
	"context"
	"fmt"
	"go_magic_probs_tool/internal/database"
	"time"
)

type BoosterRepository struct {
	db *database.DatabaseHandler
}

func NewBoosterRepository(db *database.DatabaseHandler) *BoosterRepository {
	return &BoosterRepository{db: db}
}

// FetchBoosterVariants fetches booster variants from the database
func (r *BoosterRepository) FetchBoosterVariants(ctx context.Context, setCodes, boosterNames []string) ([]BoosterVariant, error) {
	start := time.Now() // Start timing
	query := `
		SELECT c.boostername, c.boosterindex, c.setcode, c.sheetname, c.sheetpicks, w.boosterweight,
			   ROUND(w.boosterweight::NUMERIC / ws.totalweight::NUMERIC, 4) as booster_probability
		FROM setboostercontents c
		INNER JOIN setboostercontentweights w
			ON c.boostername = w.boostername AND c.setcode = w.setcode
			AND (c.boosterindex = w.boosterindex OR (c.boosterindex IS NULL AND w.boosterindex IS NULL))
		INNER JOIN (
			SELECT boostername, setcode, SUM(boosterweight) AS totalweight
			FROM setboostercontentweights
			WHERE 1 = 1
	`
	queryParams := []any{}
	paramCounter := 1

	if len(setCodes) > 0 {
		query += fmt.Sprintf(" AND setcode = ANY($%d)", paramCounter)
		queryParams = append(queryParams, setCodes)
		paramCounter++
	}

	if len(boosterNames) > 0 {
		query += fmt.Sprintf(" AND boostername = ANY($%d)", paramCounter)
		queryParams = append(queryParams, boosterNames)
		paramCounter++
	}

	query += " GROUP BY boostername, setcode ) ws ON c.boostername = ws.boostername AND c.setcode = ws.setcode"
	rows, err := r.db.Pool.Query(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var variants []BoosterVariant
	for rows.Next() {
		var variant BoosterVariant
		if err := rows.Scan(
			&variant.BoosterName,
			&variant.BoosterIndex,
			&variant.SetCode,
			&variant.SheetName,
			&variant.SheetPicks,
			&variant.BoosterWeight,
			&variant.BoosterProbability,
		); err != nil {
			return nil, err
		}
		variants = append(variants, variant)
	}

	fmt.Printf("FetchBoosterVariants Total Execution Time: %v\n", time.Since(start))
	return variants, nil
}
