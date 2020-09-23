package main

import (
	"fmt"
	"math/rand"
)

const (
	min = 1
	max = 6
)
var counterMap map[int]int

func main() {
	counterMap = make(map[int]int, max)
	for i := 1; i < 1001; i++ {
		counterMap[rand.Intn(max - min) + min]++
	}
	fmt.Printf("Counter map: %+v", counterMap)
}
