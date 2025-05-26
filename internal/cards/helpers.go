package cards

// extractCardUUIDs extracts all UUIDs from card data map
func ExtractCardUUIDs(cardData map[string]CardData) []string {
	uuids := make([]string, 0, len(cardData))
	for uuid := range cardData {
		uuids = append(uuids, uuid)
	}
	return uuids
}
