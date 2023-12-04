package round

import (
	"math"
)

func RoundP(flo float64, pre uint) float64 {
	ratio := math.Pow(10, float64(pre))
	return math.Round(flo*ratio) / ratio
}

func RoundN(f float64, n uint) float64 {
	p := math.Pow10(int(n))
	return RoundP(f/p, 0) * p
}
