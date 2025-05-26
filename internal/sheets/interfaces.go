package sheets

import (
	"context"
)

type SheetFetcher interface {
	FetchBoosterSheets(ctx context.Context, setCodes []string, boosterNames []string, cardUUIDs []string) (map[string]map[string][]BoosterSheetEntry, error)
}
