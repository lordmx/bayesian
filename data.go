package bayesian

type Data struct {
	F map[Word]int
	Total int
}

func newData() *Data {
	return &Data{
		F: make(map[string]int),
	}
}

func (data *Data) Prob(words ...Word) float64 {
	prob := 1.0

	for _, word := range words {
		value, present := data.F[word]

		if !present {
			value = 0.0000000001
		}

		prob *= float64(value) / float64(data.Total) 
	}

	return prob
}