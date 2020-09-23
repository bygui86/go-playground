package main

import (
	"fmt"
	"time"
)

const (
	timeFormat = time.RFC3339
	startStr   = "2020-06-15T23:00:00Z"
	endStr     = "2020-06-16T00:00:00Z"
)

func main() {
	start, startErr := parseTime(timeFormat, startStr)
	if startErr != nil {
		panic(startErr)
	}
	end, endErr := parseTime(timeFormat, endStr)
	if endErr != nil {
		panic(endErr)
	}

	fmt.Printf("Start: %+v \n", start)
	fmt.Printf("End: %+v \n", end)
	fmt.Printf("Delta in min: %f \n", delta(start, end, time.Minute))
	fmt.Printf("Delta in sec: %f \n", delta(start, end, time.Second))
}

func delta(start, end time.Time, format time.Duration) float64 {
	switch format {
	case time.Nanosecond:
		return float64(end.Sub(start).Nanoseconds())
	case time.Microsecond:
		return float64(end.Sub(start).Microseconds())
	case time.Millisecond:
		return float64(end.Sub(start).Milliseconds())
	case time.Second:
		return end.Sub(start).Seconds()
	case time.Minute:
		return end.Sub(start).Minutes()
	case time.Hour:
		return end.Sub(start).Hours()
	}
	return end.Sub(start).Seconds()
}

func parseTime(format, timeStr string) (time.Time, error) {
	return time.Parse(format, timeStr)
}
