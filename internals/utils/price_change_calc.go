package utils

func DailyPriceChange(currentPrice, DayBeforePrice float64) float64 {

	dailyChange := ((currentPrice - DayBeforePrice) / DayBeforePrice) * 100

	return dailyChange
}
