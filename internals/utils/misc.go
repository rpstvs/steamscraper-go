package utils

import "time"

func ConvertDate() time.Time {
	currentTime := time.Now()

	currentDate := currentTime.Truncate(24 * time.Hour)

	return currentDate
}
