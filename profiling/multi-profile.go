package main

import (
	"fmt"
	"time"

	profile "github.com/bygui86/multi-profile/v2"
)

func main() {
	fmt.Println("Profiling example")

	defer profile.
		CPUProfile(&profile.Config{}).
		Start().
		Stop()

	counter := 0
	for {
		fmt.Printf("Counting: %d\n", counter)
		counter++
		time.Sleep(1 * time.Second)
	}
}
