package main

import (
	"fmt"
	"time"
)

func main() {
	timeMeasure := StartTimeMeasure()

	time.Sleep(5 * time.Second)

	timeMeasure.StopTimeMeasure()

	fmt.Printf("start: %v\n", timeMeasure.GetStart())
	fmt.Printf("stop: %v\n", timeMeasure.GetStop())
	fmt.Printf("delta (sec): %v\n", timeMeasure.GetDelta())
	fmt.Printf("delta (m-sec): %v\n", timeMeasure.GetDelta().Milliseconds())
	fmt.Printf("delta (n-sec): %v\n", timeMeasure.GetDelta().Nanoseconds())
}

type TimeMeasure struct {
	start time.Time
	stop  time.Time
	delta time.Duration
}

func StartTimeMeasure() *TimeMeasure {
	return &TimeMeasure{
		start: time.Now(),
		stop:  time.Time{},
		delta: -1,
	}
}

func (t *TimeMeasure) StopTimeMeasure() {
	t.stop = time.Now()
	t.delta = t.stop.Sub(t.start)
}

// ACCESSORS

func (t *TimeMeasure) GetStart() time.Time {
	return t.start
}

func (t *TimeMeasure) GetStop() time.Time {
	return t.stop
}

func (t *TimeMeasure) GetDelta() time.Duration {
	return t.delta
}
