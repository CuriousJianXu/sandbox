package utils

import (
	"math/rand"
	"time"
)

// GetLocalTimeInTaipei returns the local time in "Asia/Taipei" timezone for a specified
// number of days before the current date, formatted as "yyyy/mm/dd"
func GetLocalTimeInTaipei(daysBefore int) (string, error) {
	// Set the timezone to "Asia/Taipei" (GMT+8)
	taipeiLocation, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return "", err
	}

	// Get the current local time in "Asia/Taipei" timezone
	localTime := time.Now().AddDate(0, 0, -daysBefore).In(taipeiLocation)

	// Format the local time as "yyyy/mm/dd"
	formattedTime := localTime.Format("2006/01/02")

	return formattedTime, nil
}

func getRandomDuration() int {
	// Generate a random float64 between min and max
	min := 500
	max := 2000
	// nolint:gosec
	return rand.Intn(max-min) + min
}

func FriendlyCrawl() {
	duration := getRandomDuration()
	time.Sleep(time.Duration(duration) * time.Millisecond)
}
