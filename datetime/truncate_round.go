package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Printf("now: %v\n", now)

	roundSec := now.Round(time.Second)
	roundMin := now.Round(time.Minute)

	fmt.Printf("round second: %v\n", roundSec)
	fmt.Printf("round minute: %v\n", roundMin)

	truncSec := now.Truncate(time.Second)
	truncMin := now.Truncate(time.Minute)

	fmt.Printf("truncate second: %v\n", truncSec)
	fmt.Printf("truncate minute: %v\n", truncMin)
}
