package sheets

import (
	"context"
)

// SheetService handles sheet-related operations
type SheetService struct {
	Fetcher SheetFetcher
}

// NewSheetService creates a new SheetService with a fetcher
func NewSheetService(fetcher SheetFetcher) *SheetService {
	return &SheetService{
		Fetcher: fetcher,
	}
}

// GetBoosterSheets retrieves sheets using the fetcher
func (s *SheetService) FetchBoosterSheets(ctx context.Context, setCodes []string, boosterNames []string, cardUuids []string) (map[string]map[string][]BoosterSheetEntry, error) {
	return s.Fetcher.FetchBoosterSheets(ctx, setCodes, boosterNames, cardUuids)
}
