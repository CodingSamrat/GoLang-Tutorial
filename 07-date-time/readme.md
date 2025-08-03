# Working with Date & Time in Go
Go provides a powerful `time` package to work with *dates* and *times*.

#### Example
```go
fmt.Println("Welcome to Go Tutorial - DateTime!")

// Get the current date-time
currentTime := time.Now()

// Print full datetime
fmt.Println("Current timestamp =>", currentTime)

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
```

#### ⚠️Notes:
- Go uses a reference time: `Mon Jan 2 15:04:05 MST 2006`.
- You must format dates using this layout - it’s not a typo!
- `time.Now()` gives current local time.
- `Format()` is used to display the *date/time* in desired formats.

<br/>

> ***N.B.:*** For advanced usage of time, including time zones, durations, and parsing layouts, refer to the official Go documentation: [pkg.go.dev/time](https://pkg.go.dev/time)