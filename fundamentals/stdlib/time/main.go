package main

import (
    "fmt"
    "time"
)

func main() {
    // ----- Current time
    now := time.Now()
    fmt.Println("Current time:", now)

    // ----- Formatting time
    // Go uses a reference time: Mon Jan 2 15:04:05 MST 2006
    // You format by writing how you'd show that specific moment
    fmt.Println("Date:", now.Format("2006-01-02"))
    fmt.Println("Time:", now.Format("15:04:05"))
    fmt.Println("DateTime:", now.Format("2006-01-02 15:04:05"))
    fmt.Println("Custom:", now.Format("Mon, 02 Jan 2006"))

    // Predefined format constants
    fmt.Println("RFC3339:", now.Format(time.RFC3339))
    fmt.Println("Kitchen:", now.Format(time.Kitchen))

    // ----- Parsing time from string
    // The layout must use the same reference time
    layout := "2006-01-02"
    parsed, err := time.Parse(layout, "2025-12-25")
    if err != nil {
        fmt.Println("Parse error:", err)
        return
    }
    fmt.Println("Parsed date:", parsed)

    // ----- Creating specific dates
    birthday := time.Date(2000, time.January, 15, 10, 30, 0, 0, time.UTC)
    fmt.Println("Birthday:", birthday)

    // ----- Time components
    fmt.Println("Year:", now.Year())
    fmt.Println("Month:", now.Month())
    fmt.Println("Day:", now.Day())
    fmt.Println("Hour:", now.Hour())
    fmt.Println("Minute:", now.Minute())
    fmt.Println("Weekday:", now.Weekday())

    // ----- Durations
    // time.Duration represents a length of time
    duration := 2*time.Hour + 30*time.Minute + 15*time.Second
    fmt.Println("Duration:", duration)           // 2h30m15s
    fmt.Println("In seconds:", duration.Seconds()) // 9015
    fmt.Println("In minutes:", duration.Minutes()) // 150.25

    // ----- Adding and subtracting time
    tomorrow := now.Add(24 * time.Hour)
    fmt.Println("Tomorrow:", tomorrow.Format("2006-01-02"))

    yesterday := now.Add(-24 * time.Hour)
    fmt.Println("Yesterday:", yesterday.Format("2006-01-02"))

    // AddDate: add years, months, days
    nextYear := now.AddDate(1, 0, 0)
    fmt.Println("Next year:", nextYear.Format("2006-01-02"))

    // ----- Comparing times
    fmt.Println("Is tomorrow after now?", tomorrow.After(now))   // true
    fmt.Println("Is yesterday before now?", yesterday.Before(now)) // true

    // ----- Time difference
    diff := tomorrow.Sub(now)
    fmt.Println("Difference:", diff) // 24h0m0s

    // ----- Timers and sleeping (commented to avoid blocking)
    // time.Sleep(2 * time.Second) // pauses execution for 2 seconds
    // timer := time.NewTimer(5 * time.Second)
    // <-timer.C // blocks until the timer fires

    // ----- Tickers (commented to avoid blocking)
    // ticker := time.NewTicker(1 * time.Second)
    // defer ticker.Stop()
    // for i := 0; i < 3; i++ {
    //     <-ticker.C
    //     fmt.Println("Tick!", i)
    // }

    // ----- Unix timestamps
    unix := now.Unix()
    fmt.Println("Unix timestamp:", unix)
    fromUnix := time.Unix(unix, 0)
    fmt.Println("From Unix:", fromUnix.Format("2006-01-02 15:04:05"))
}
