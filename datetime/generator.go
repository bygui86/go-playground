package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	counter := 0
	for y := 0; y < 3; y++ { // YEAR
		year := now.AddDate(-y, 0, 0)

		for m := 0; m < 12; m++ { // MONTH
			// for m := 0; m < 3; m++ { // MONTH
			month := year.AddDate(0, -m, 0)

			for d := 0; d < 30; d++ { // DAY
				// for d := 0; d < 3; d++ { // DAY
				day := month.AddDate(0, 0, -d)

				fmt.Printf("%s\n", day.Format(time.RFC3339))
				counter++
			}
		}
	}

	fmt.Printf("Total date generated: %d\n", counter)
}
