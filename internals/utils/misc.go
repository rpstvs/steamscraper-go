package utils

import (
	"strings"
	"time"
)

func ConvertDate() time.Time {
	currentTime := time.Now()

	// Truncate to get only the date part
	currentDate := currentTime.Truncate(24 * time.Hour)

	return currentDate
}

func ExtractSteamid(steamid string) string {
	tmp := strings.Split(steamid, "/")
	id := tmp[len(tmp)-1]
	return id
}
