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

type DrilldownCard struct {
	UUID         string   `json:"uuid"`
	Foil         bool     `json:"foil"`
	PromoTypes   []string `json:"promoTypes,omitempty"`
	FrameEffects []string `json:"frameEffects,omitempty"`
	Probability  float64  `json:"probability"`
}

type DrilldownBooster struct {
	BoosterName string          `json:"boosterName"`
	Cards       []DrilldownCard `json:"cards"`
	TotalProb   float64         `json:"totalProbability"`
}

type DrilldownSet struct {
	SetCode   string             `json:"setCode"`
	Boosters  []DrilldownBooster `json:"boosters"`
	TotalProb float64            `json:"totalProbability"`
}

type DrilldownResponse struct {
	Sets []DrilldownSet `json:"sets"`
}

type CalculateResponse struct {
	Probabilities       map[string][]CardProbability `json:"probabilities"`
	AggregatedByBooster map[string]float64           `json:"aggregatedByBooster"`
	AggregatedByFoil    map[string]float64           `json:"aggregatedByFoil"`
	AggregatedBySet     map[string]float64           `json:"aggregatedBySet"`
	Drilldown           []DrilldownSet               `json:"drilldown"`
	Error               string                       `json:"error,omitempty"`
}

type Row struct {
	UUID           string
	Booster        string
	BoosterVariant *int
	Set            string
	Foil           bool
	PromoTypes     []string
	FrameEffects   []string
	Sheet          string
	SheetPicks     int
	Probability    float64
}
