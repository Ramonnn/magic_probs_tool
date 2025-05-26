package calculate

func AggregateCardProbabilities(
	raw map[string]map[string][]CardProbability,
) map[string]map[string][]CardProbability {
	aggregated := make(map[string]map[string][]CardProbability)

	for boosterName, cardsMap := range raw {
		if _, exists := aggregated[boosterName]; !exists {
			aggregated[boosterName] = make(map[string][]CardProbability)
		}

		for cardID, probList := range cardsMap {
			// Use a map key to merge CardProbability by IsFoil + SetCode
			mergeMap := make(map[string]CardProbability)

			for _, cp := range probList {
				// Create a unique key per foil + setcode to aggregate those entries
				key := cp.SetCode + "|" + boolToStr(cp.IsFoil)

				if existing, found := mergeMap[key]; found {
					existing.Probability += cp.Probability
					mergeMap[key] = existing
				} else {
					mergeMap[key] = cp
				}
			}

			// Convert merged map back to slice
			mergedSlice := make([]CardProbability, 0, len(mergeMap))
			for _, val := range mergeMap {
				mergedSlice = append(mergedSlice, val)
			}

			aggregated[boosterName][cardID] = mergedSlice
		}
	}

	return aggregated
}

// Helper to convert bool to string
func boolToStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
