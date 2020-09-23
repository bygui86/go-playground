package main

import (
	"fmt"
	"math/big"
)

func main() {
	examples()

	// values := []float64{1000000000000000.101010101010101010101001, 1000000000001000.101010101010101010101001}
	// floatCalculation(values)
	// mathBigCalculation(values)
}

func examples() {
	// Operate on numbers of different precision.
	var x, y, z big.Float
	x.SetInt64(1000)          // x is automatically set to 64bit precision
	y.SetFloat64(2.718281828) // y is automatically set to 53bit precision
	z.SetPrec(32)
	// z.SetPrec(64)
	// z.SetPrec(128)
	z.Add(&x, &y)
	fmt.Printf("x = %.10g (%s, prec = %d, acc = %s)\n", &x, x.Text('p', 0), x.Prec(), x.Acc())
	fmt.Printf("y = %.10g (%s, prec = %d, acc = %s)\n", &y, y.Text('p', 0), y.Prec(), y.Acc())
	fmt.Printf("z = %.13g (%s, prec = %d, acc = %s)\n", &z, z.Text('p', 0), z.Prec(), z.Acc())

	if z.Acc() == big.Below {
		z.SetPrec(64)
		z.Add(&x, &y)
		fmt.Printf("z = %.13g (%s, prec = %d, acc = %s)\n", &z, z.Text('p', 0), z.Prec(), z.Acc())
	}

	fmt.Printf("z = %f\n", &z)
}

func floatCalculation(values []float64) {
	sum := 0.0
	for _, x := range values {
		sum = sum + x
	}
	fmt.Printf("Standard calculation (float): %f \n", sum)
}

func mathBigCalculation(values []float64) {
	var sum big.Float
	sum.SetPrec(64)
	// sum.SetPrec(128)
	for _, x := range values {
		sum.Add(&sum, big.NewFloat(x))
	}
	// fmt.Printf("z = %.10g (%s, prec = %d, acc = %s)\n", &z, z.Text('p', 0), z.Prec(), z.Acc())
	fmt.Printf("Decimal calculation: (float) %f  /  (string) %s \n", &sum, sum.String())
}
