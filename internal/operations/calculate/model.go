package calculate

type CardProbability struct {
	Probability    float64  `json:"probability"`
	IsFoil         bool     `json:"isFoil"`
	SetCode        string   `json:"setCode"`
	BoosterName    string   `json:"boosterName"`
	PromoTypes     []string `json:"promoTypes"`
	FrameEffects   []string `json:"frameEffects"`
	SheetName      string   `json:"sheetName"`
	SheetPicks     int      `json:"sheetPicks"`
	BoosterVariant *int     `json:"boosterVariant"`
}

type CalculateRequest struct {
	Cards []string `json:"cards"`
}

type CalculateResponse struct {
	Probabilities       map[string][]CardProbability `json:"probabilities"`
	AggregatedByBooster map[string]float64           `json:"aggregatedByBooster"`
	AggregatedByFoil    map[string]float64           `json:"aggregatedByFoil"`
	Error               string                       `json:"error,omitempty"`
}

type Row struct {
	Booster        string
	BoosterVariant *int
	UUID           string
	Sheet          string
	SheetPicks     int
	Set            string
	Foil           bool
	PromoTypes     []string
	FrameEffects   []string
	Probability    float64
}
