package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	// values := [4]float64{0.032367, 0.009490, 0.015570, 0.031024}
	// total 0.08845085999999999

	// values := [2]float64{0.006790, 0.100000}
	// total 0.10679000000000001

	// values := [41]float64{0.146350, 0.004000, 0.019786, 0.022559, 0.024909, 0.024909, 0.024909, 0.022565, 0.002343, 0.014797, 0.010112, 0.024909, 0.024909, 0.024909, 0.024909, 0.024909, 0.042345, 0.024909, 0.024909, 0.024909, 0.024909, 0.042345, 0.042345, 0.024909, 0.024909, 0.042345, 0.042345, 0.042345, 0.024909, 0.024909, 0.024909, 0.042345, 0.024909, 0.024909, 0.266004, 0.177140, 0.068070, 0.174200, 0.051710, 0.057490, 0.132990}
	// total 1.9397942300000022

	// values := []float64{1000000000000.101010101010101010101001, 1000000000001.101010101010101010101001}
	// total 2000000000001.201904
	values := []float64{1000000000000000.101010101010101010101001, 1000000000001000.101010101010101010101001}

	standardCalculation2(values)

	decimalCalculation2(values)
}

func standardCalculation2(values []float64) {
	sum := 0.0
	for _, x := range values {
		sum = sum + x
	}
	fmt.Printf("Standard calculation (float): %f \n", sum)
}

func decimalCalculation2(values []float64) {
	sumDec := decimal.NewFromFloat(0.0)
	for _, x := range values {
		sumDec = sumDec.Add(decimal.NewFromFloat(x))
	}
	sum, _ := sumDec.Float64()
	fmt.Printf("Decimal calculation: (float) %f  /  (string) %s \n", sum, sumDec.String())
}
