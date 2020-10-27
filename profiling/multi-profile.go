package main

import (
	"fmt"
	"time"

	"github.com/bygui86/multi-profile"
)

func main() {
	fmt.Println("Profiling example")

	defer profile.
		CPUProfile(&profile.ProfileConfig{}).
		Start().
		Stop()

	counter := 0
	for {
		fmt.Printf("Counting: %d\n", counter)
		counter++
		time.Sleep(1 * time.Second)
	}
}
