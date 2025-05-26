package sheets

import (
	"context"
	"database/sql"
	"fmt"
	"go_magic_probs_tool/internal/database"
	"strings"
	"time"
)

type SheetRepository struct {
	db *database.DatabaseHandler
}

func NewSheetRepository(db *database.DatabaseHandler) *SheetRepository {
	return &SheetRepository{db: db}
}

func (r *SheetRepository) FetchBoosterSheets(ctx context.Context, setCodes, boosterNames, cardUUIDs []string) (map[string]map[string][]BoosterSheetEntry, error) {

	start := time.Now() // Start timing
	baseQuery := `
        SELECT
		c.carduuid, 
		c.boostername, 
		c.setcode, 
		c.sheetname, 
		c.cardweight, 
		s.sheetweight,
		ROUND(c.cardweight::NUMERIC / s.sheetweight::NUMERIC, 4) AS card_probability,
               d.sheetisfoil
        FROM setboostersheetcards c
        LEFT JOIN setboostersheets d
          ON c.boostername = d.boostername 
          AND c.sheetname = d.sheetname
          AND c.setcode = d.setcode
        INNER JOIN (
            SELECT boostername, setcode, sheetname, SUM(cardweight) AS sheetweight
            FROM setboostersheetcards
            WHERE 1 = 1
    `

	var whereClauses []string
	var args []any
	argPos := 1

	if len(setCodes) > 0 {
		placeholders := make([]string, len(setCodes))
		for i, code := range setCodes {
			placeholders[i] = fmt.Sprintf("$%d", argPos)
			args = append(args, code)
			argPos++
		}
		whereClauses = append(whereClauses, fmt.Sprintf("AND setcode IN (%s)", strings.Join(placeholders, ", ")))
	}

	if len(boosterNames) > 0 {
		placeholders := make([]string, len(boosterNames))
		for i, name := range boosterNames {
			placeholders[i] = fmt.Sprintf("$%d", argPos)
			args = append(args, name)
			argPos++
		}
		whereClauses = append(whereClauses, fmt.Sprintf("AND boostername IN (%s)", strings.Join(placeholders, ", ")))
	}

	baseQuery += strings.Join(whereClauses, " ") + `
            GROUP BY boostername, setcode, sheetname
        ) s
          ON c.boostername = s.boostername AND c.setcode = s.setcode AND c.sheetname = s.sheetname
        WHERE 1=1
    `

	if len(cardUUIDs) > 0 {
		placeholders := make([]string, len(cardUUIDs))
		for i, uuid := range cardUUIDs {
			placeholders[i] = fmt.Sprintf("$%d", argPos)
			args = append(args, uuid)
			argPos++
		}
		baseQuery += fmt.Sprintf(" AND c.carduuid IN (%s)", strings.Join(placeholders, ", "))
	}

	// Initialize nested map: map[sheetName]map[cardUUID][]BoosterSheetEntry
	sheets := make(map[string]map[string][]BoosterSheetEntry)

	queryStart := time.Now()
	rows, err := r.db.Pool.Query(ctx, baseQuery, args...)
	fmt.Printf("Query FetchBoosterSheets executed in: %v\n", time.Since(queryStart))
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var isFoil sql.NullBool
	for rows.Next() {
		var entry BoosterSheetEntry
		err := rows.Scan(
			&entry.CardUUID,
			&entry.BoosterName,
			&entry.SetCode,
			&entry.SheetName,
			&entry.CardWeight,
			&entry.SheetWeight,
			&entry.CardProbability,
			&isFoil,
		)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}

		if isFoil.Valid {
			entry.IsFoil = isFoil.Bool
		} else {
			entry.IsFoil = false
		}

		if _, ok := sheets[entry.SheetName]; !ok {
			sheets[entry.SheetName] = make(map[string][]BoosterSheetEntry)
		}
		sheets[entry.SheetName][entry.CardUUID] = append(sheets[entry.SheetName][entry.CardUUID], entry)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	fmt.Printf("FetchBoosterSheets Total Execution Time: %v\n", time.Since(start))
	return sheets, nil
}
