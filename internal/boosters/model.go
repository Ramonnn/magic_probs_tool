package boosters

// BoosterVariant represents a booster variant and its associated probabilities.
type BoosterVariant struct {
	BoosterName        string
	BoosterIndex       *int
	SetCode            string
	SheetName          string
	SheetPicks         int
	BoosterWeight      float64
	BoosterProbability float64
}
