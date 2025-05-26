package sheets

// BoosterSheetEntry represents a card's data on a booster sheet.
type BoosterSheetEntry struct {
	CardUUID        string
	BoosterName     string
	SetCode         string
	SheetName       string
	CardWeight      float64
	SheetWeight     float64
	CardProbability float64
	IsFoil          bool
}
