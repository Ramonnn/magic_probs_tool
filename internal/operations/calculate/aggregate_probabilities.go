package calculate

func AggregateBy(rows []Row, keyFunc func(r Row) string) map[string]float64 {
	agg := make(map[string]float64)
	for _, r := range rows {
		key := keyFunc(r)
		agg[key] += r.Probability
	}
	return agg
}
