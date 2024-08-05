package utils

func DailyPriceChange(currentPrice, DayBeforePrice float64) float64 {

	dailyChange := ((currentPrice - DayBeforePrice) / DayBeforePrice) * 100

	return dailyChange
}

func WeeklyPriceChange(prices []float64) float64 {

	var sum float64
	var average float64

	for x := range prices {
		sum += prices[x]
	}

	average = sum / 7.0

	WeeklyChange := ((prices[0] - average) / average) * 100

	return WeeklyChange
}
