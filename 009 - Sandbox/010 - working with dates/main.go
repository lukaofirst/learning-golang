package main

import (
	"fmt"
	"time"
)

func main() {
	// Get current time
	currentTime := time.Now()
	fmt.Println(currentTime)
	
	// Get current time as UTC
	currentTimeUTC := currentTime.UTC()
	fmt.Println(currentTimeUTC)

	// Formatting current time UTC to string UTC format
	utcString := currentTimeUTC.Format(time.RFC3339)
	fmt.Println(utcString)

	// Adding hours
	currentTimePlusHours := currentTimeUTC.Add(2 * time.Hour).Format(time.RFC3339)
	fmt.Println(currentTimePlusHours)

	// Adding minutes
	currentTimePlusMinutes := currentTimeUTC.Add(30 * time.Minute).Format(time.RFC3339)
	fmt.Println(currentTimePlusMinutes)
	fmt.Println(currentTimePlusHours)

	// Adding seconds
	currentTimePlusSeconds := currentTimeUTC.Add(15 * time.Second).Format(time.RFC3339)
	fmt.Println(currentTimePlusSeconds)

	// Adding seconds
	currentTimePlusMilliseconds := currentTimeUTC.Add(20000 * time.Millisecond).Format(time.RFC3339)
	fmt.Println(currentTimePlusMilliseconds)
	
	// Converting date string into date
	layout := "2006-01-02 15:04:05"
	utcString2 := "2024-08-16 12:00:00"
	parsedTime, err := time.Parse(layout, utcString2)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	} else {
		fmt.Println("Parsed time:", parsedTime)
	}
}