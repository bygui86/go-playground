package main

import (
	"fmt"
	"math"

	"github.com/shopspring/decimal"
)

const decimalCiphers = 6

var values = [4]float64{0.032367, 0.009490, 0.015570, 0.031024}

func main() {
	// mathExamples()

	standardCalculation()

	// roundedStandardCalculation()

	// decimalCalculation()

	// roundedDecimalCalculation()
}

func mathExamples() {
	u := math.RoundToEven(11.5)
	fmt.Printf("%.1f\n", u)

	d := math.RoundToEven(12.5)
	fmt.Printf("%.1f\n", d)

	p := math.Round(10.4)
	fmt.Printf("%.1f\n", p)

	n := math.Round(-10.5)
	fmt.Printf("%.1f\n", n)
}

func standardCalculation() {
	sum := 0.0
	for _, x := range values {
		sum = sum + x
	}
	fmt.Printf("Standard calculation (value): %v \n", sum)
	fmt.Printf("Standard calculation (float): %f \n", sum)
}

func roundedStandardCalculation() {
	sum := 0.0
	for _, x := range values {
		sum = sum + x
	}
	fmt.Printf("Rounded standard calculation: %f \n", math.RoundToEven(sum))
}

func decimalCalculation() {
	sumDec := decimal.NewFromFloat(0.0)
	for _, x := range values {
		sumDec = sumDec.Add(decimal.NewFromFloat(x))
	}
	sum, _ := sumDec.Float64()
	fmt.Printf("Decimal calculation: (float) %f  /  (string) %s \n", sum, sumDec.String())
}

func roundedDecimalCalculation() {
	sumDec := decimal.NewFromFloat(0.0)
	for _, x := range values {
		sumDec = sumDec.Add(decimal.NewFromFloat(x))
	}
	sumDec = sumDec.Round(decimalCiphers)
	sum, _ := sumDec.Float64()
	fmt.Printf("Rounded decimal calculation: (float) %f  /  (string) %s \n", sum, sumDec.String())
}
