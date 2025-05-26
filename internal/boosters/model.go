package boosters

// BoosterVariant represents a booster variant and its associated probabilities.
type BoosterVariant struct {
	BoosterName        string  `json:"boosterName"`
	BoosterIndex       *int    `json:"boosterIndex"`
	SetCode            string  `json:"setCode"`
	SheetName          string  `json:"sheetName"`
	SheetPicks         int     `json:"sheetPicks"`
	BoosterWeight      float64 `json:"boosterWeight"`
	BoosterProbability float64 `json:"boosterProbability"`
}
