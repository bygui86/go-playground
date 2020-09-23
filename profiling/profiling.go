package main

import (
	"fmt"
	"time"

	"github.com/pkg/profile"
)

func main() {
	fmt.Println("Profiling example")

	defer profile.
		Start(
			profile.MemProfile,
			profile.ProfilePath(".")).
		Stop()

	prof := profile.Profile{}
	profile.Start()

	counter := 0
	for {
		fmt.Printf("Counting: %d\n", counter)
		counter++
		time.Sleep(1 * time.Second)
	}
}
