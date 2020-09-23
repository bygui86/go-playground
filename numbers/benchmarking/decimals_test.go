package benchmarking_test

import (
	"math/big"
	"testing"

	"github.com/shopspring/decimal"
)

/*
	Command to run this benchmark

		go test -bench . -benchmem
*/

var values = [41]float64{0.146350, 0.004000, 0.019786, 0.022559, 0.024909, 0.024909, 0.024909, 0.022565, 0.002343,
	0.014797, 0.010112, 0.024909, 0.024909, 0.024909, 0.024909, 0.024909, 0.042345, 0.024909, 0.024909, 0.024909,
	0.024909, 0.042345, 0.042345, 0.024909, 0.024909, 0.042345, 0.042345, 0.042345, 0.024909, 0.024909, 0.024909,
	0.042345, 0.024909, 0.024909, 0.266004, 0.177140, 0.068070, 0.174200, 0.051710, 0.057490, 0.132990}

func BenchmarkShopspring(b *testing.B) {
	sumDec := decimal.NewFromFloat(0.0)
	for _, x := range values {
		sumDec = sumDec.Add(decimal.NewFromFloat(x))
	}
	sumDec.Float64()
}

func BenchmarkMathBig(b *testing.B) {
	var sumDec big.Float
	sumDec.SetPrec(64) // precision can be increased
	for _, x := range values {
		sumDec.Add(&sumDec, big.NewFloat(x))
	}
	sumDec.Float64()
}
