package calculate

type CardProbability struct {
	Probability float64 `json:"probability"`
	IsFoil      bool    `json:"isFoil"`
	SetCode     string  `json:"setCode"`
	BoosterName string  `json:"boosterName"`
}
