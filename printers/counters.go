package main

import (
	"fmt"
)

func main() {
	everyX(100)
}

func everyX(sentinel int) {
	for i := 0; i < 5000; i++ {
		if i%sentinel == 0 {
			fmt.Printf("Iteration %d\n", i)
		}
	}
}
