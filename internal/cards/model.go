package cards

type CardData struct {
	UUID         string   `json:"uuid"`
	Name         string   `json:"name"`
	Number       string   `json:"number"`
	FrameEffects []string `json:"frameEffects"`
	PromoTypes   []string `json:"promoTypes"`
}
