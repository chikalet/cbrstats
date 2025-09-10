package stats

type StatResult struct {
	MaxRate     float64
	MaxCurrency string
	MaxDate     string

	MinRate     float64
	MinCurrency string
	MinDate     string

	AverageRate float64
	TotalCount  int
}

type Aggregator struct {
	totalSum float64
	count    int

	maxRate     float64
	maxCurrency string
	maxDate     string

	minRate     float64
	minCurrency string
	minDate     string
}

func NewAggregator() *Aggregator {
	return &Aggregator{
		maxRate: -1,
		minRate: 1e308,
	}
}

func (a *Aggregator) Add(rate float64, currency, date string) {
	a.totalSum += rate
	a.count++

	if rate > a.maxRate {
		a.maxRate = rate
		a.maxCurrency = currency
		a.maxDate = date
	}
	if rate < a.minRate {
		a.minRate = rate
		a.minCurrency = currency
		a.minDate = date
	}
}

func (a *Aggregator) Result() StatResult {
	return StatResult{
		MaxRate:     a.maxRate,
		MaxCurrency: a.maxCurrency,
		MaxDate:     a.maxDate,
		MinRate:     a.minRate,
		MinCurrency: a.minCurrency,
		MinDate:     a.minDate,
		AverageRate: a.totalSum / float64(a.count),
		TotalCount:  a.count,
	}
}
