package cards

import (
	"context"
	"database/sql"
	"fmt"
	"go_magic_probs_tool/internal/database"
	"strings"
	"time"
)

type CardRepository struct {
	db *database.DatabaseHandler
}

func NewCardRepository(db *database.DatabaseHandler) *CardRepository {
	return &CardRepository{db: db}
}

// FetchCardData fetches card data for the given card names in chunks to avoid too many placeholders.
func (r *CardRepository) FetchCardData(ctx context.Context, cardNames []string, chunkSize int) (map[string]CardData, error) {
	start := time.Now() // Start timing

	if len(cardNames) == 0 {
		return map[string]CardData{}, nil
	}

	cards := make(map[string]CardData)

	for start := 0; start < len(cardNames); start += chunkSize {
		end := start + chunkSize
		end = min(len(cardNames), end)
		chunk := cardNames[start:end]

		placeholders := make([]string, len(chunk))
		args := make([]any, len(chunk))
		for i, name := range chunk {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
			args[i] = name
		}

		query := fmt.Sprintf(
			"SELECT uuid, name, number, frameeffects, promotypes FROM cards WHERE name IN (%s);",
			strings.Join(placeholders, ", "),
		)
		// Measure query time
		queryStart := time.Now()
		rows, err := r.db.Pool.Query(ctx, query, args...)
		fmt.Printf("FetchCardData Query executed in: %v\n", time.Since(queryStart))
		if err != nil {
			return nil, fmt.Errorf("query failed: %w", err)
		}

		for rows.Next() {
			var c CardData
			var frameEffects sql.NullString
			var promoTypes sql.NullString

			err := rows.Scan(&c.UUID, &c.Name, &c.Number, &frameEffects, &promoTypes)
			if err != nil {
				rows.Close()
				return nil, fmt.Errorf("scan failed: %w", err)
			}

			if frameEffects.Valid && frameEffects.String != "" {
				parts := strings.Split(frameEffects.String, ",")
				c.FrameEffects = &parts
			} else {
				c.FrameEffects = &[]string{}
			}

			if promoTypes.Valid && promoTypes.String != "" {
				parts := strings.Split(promoTypes.String, ",")
				c.PromoTypes = &parts
			} else {
				c.PromoTypes = &[]string{}
			}

			cards[c.UUID] = c
		}

		rows.Close()
		if rows.Err() != nil {
			return nil, rows.Err()
		}
	}
	fmt.Printf("FetchCardData Total Execution Time: %v\n", time.Since(start))
	return cards, nil
}
