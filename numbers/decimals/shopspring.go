package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {

	// DivisionPrecision is the number of decimal places in the result when it doesn't divide exactly.
	// decimal.DivisionPrecision = 16 // default 16

	floatCalc(165.000000000000000000000000, 1.400000000000000000000000)
	floatCalc(165.000000000000000000000000, 1.400000000000000000000000)

	f := 0.020000000000000018000000
	printFloat(f)

	decimalMultCalc(165.0, 1.4)

	decimalDivCalc(2, 3)
	decimalDivCalc(2, 30000)
	decimalDivCalc(20000, 3)

	printDecimalFromFloat(f)
}

func floatCalc(x, y float64) {
	result := x * y
	fmt.Printf("%.24f * %.24f = %.24f\n", x, y, result)
}

func decimalMultCalc(x, y float64) {
	result := decimal.NewFromFloat(x).Mul(decimal.NewFromFloat(y))
	fmt.Printf("%.16f * %.16f = %s\n", x, y, result.String())
}

func decimalDivCalc(x, y float64) {
	result := decimal.NewFromFloat(x).Div(decimal.NewFromFloat(y))
	fmt.Printf("%.16f / %.16f = %s\n", x, y, result.String())
}

func printFloat(f float64) {
	fmt.Printf("%.24f\n", f)
}

func printDecimal(d decimal.Decimal) {
	fmt.Printf("%s\n", d.String())
}

func printDecimalFromFloat(f float64) {
	d := decimal.NewFromFloat(f)
	fmt.Printf("%s\n", d.String())
}
