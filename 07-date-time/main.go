package main

import (
	"fmt"
	"time"
)

func main() {
fmt.Println("Welcome to Go Tutorial - DateTime!")

// Get the current date-time
currentTime := time.Now()

// Print full datetime
fmt.Println("Current Date & Time =>", currentTime)

// Print only date
fmt.Println("Date YYYY-MM-DD =>", currentTime.Format("2006-01-02"))
fmt.Println("Date MM-DD-YYYY =>", currentTime.Format("01-02-2006"))
fmt.Println("Date DD-MM-YYYY =>", currentTime.Format("02-01-2006"))

// Print only time
fmt.Println("Time HH:MM:SS =>", currentTime.Format("15:04:05"))
fmt.Println("Time HH:MM =>", currentTime.Format("15:04"))

// Date time
fmt.Println("Date Time DD-MM-YYYY HH:MM =>", currentTime.Format("02-01-2006 15:04"))

// Date time and day
fmt.Println("Date Time DD-MM-YYYY HH:MM =>", currentTime.Format("02-01-2006 15:04 Monday"))


}