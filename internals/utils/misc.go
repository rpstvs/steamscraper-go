package utils

import "time"

func ConvertDate() time.Time {
	currentTime := time.Now()

	// Truncate to get only the date part
	currentDate := currentTime.Truncate(24 * time.Hour)

	return currentDate
}
