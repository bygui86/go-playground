package main

import (
	"fmt"
	"time"
)

// see https://golang.org/src/time/example_test.go

func main() {
	// printUnix()
	// fromUnixToTime()
	// unixCalc()
	// units()
	printFormatDateTime()
}

// WARN: running on macOs the precision could be limited by the OS to microseconds
func printFormatDateTime() {
	now := time.Now()
	nowNanos := time.Now().UnixNano()

	fmt.Printf("%v \n", now)
	fmt.Printf("%+v \n", now)

	// time.Time's Stringer method is useful without any format.
	fmt.Println("default format:", now)

	// Predefined constants in the package implement common layouts.
	fmt.Println("Unix format:", now.Format(time.UnixDate))

	// The time zone attached to the time value affects its output.
	fmt.Println("Same, in UTC:", now.UTC().Format(time.UnixDate))

	// The time zone attached to the time value affects its output.
	fmt.Println("Time to nanos:", now.Format("15:04:05.000000000"))
	fmt.Println("Time to nanos:", now.Format(time.StampNano))

	fmt.Println("Time to nanos:", time.Unix(0, nowNanos).Format("15:04:05.000000000"))
	fmt.Println("Time to nanos:", time.Unix(0, nowNanos).Format(time.StampNano))

	a := time.Now()
	b := time.Now()
	fmt.Println(b.Sub(a))
}

func units() {
	now := time.Now()
	nowUnix := now.Unix()         // seconds
	nowUnixNano := now.UnixNano() // nanos

	nowBackFromUnix := time.Unix(nowUnix, 0)
	nowBackFromUnixNano := time.Unix(0, nowUnixNano)

	fmt.Printf("now: %v \n", now)
	fmt.Printf("now unix: %d \n", nowUnix)
	fmt.Printf("now unix nano: %d \n", nowUnixNano)
	fmt.Printf("now back from unix: %v \n", nowBackFromUnix)
	fmt.Printf("now back from unix nano: %v \n", nowBackFromUnixNano)
}

func unixCalc() {
	nowUnix := time.Now().UTC().Unix()
	endUnix := nowUnix - 121
	startUnix := nowUnix

	endTime := time.Unix(endUnix, 0)
	startTime := time.Unix(startUnix, 0)
	fmt.Printf("start: %v / %d\n", startTime, startUnix)
	fmt.Printf("end: %v / %d\n", endTime, endUnix)
}

func fromUnixToTime() {
	// unixTime := time.Now().Unix()
	var unixTime int64 = 1573142098
	humanTime := time.Unix(unixTime, 0) // time.Unix(seconds, nanoseconds)
	strDate := humanTime.Format(time.UnixDate)
	fmt.Println(strDate)
}

func printUnix() {
	fmt.Printf("time now unix: %d\n", time.Now().Unix())
}
