package sheets

func ExtractSetCodesFromSheets(sheets map[string]map[string][]BoosterSheetEntry) []string {
	setCodeSet := make(map[string]struct{})
	for _, cardMap := range sheets {
		for _, entries := range cardMap {
			if len(entries) > 0 {
				setCodeSet[entries[0].SetCode] = struct{}{}
			}
		}
	}

	setCodes := make([]string, 0, len(setCodeSet))
	for setCode := range setCodeSet {
		setCodes = append(setCodes, setCode)
	}
	return setCodes
}
