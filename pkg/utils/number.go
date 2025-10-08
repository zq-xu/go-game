package utils

import (
	"math/rand"

	"github.com/samber/lo"
)

func GetMinFloat64(a, b float64) float64 {
	if a < b {
		return a
	}

	return b
}

// RandomForMinFloat64
func RandomForMinFloat64(a, b float64) float64 {
	return lo.Min([]float64{a, b}) * rand.Float64()
}

// AbsPlusOneKeepSign
func AbsPlusOneKeepSign(x float64) float64 {
	return lo.Ternary(x == 0, 0.0, lo.Ternary(x > 0, x+1, -(-x+1)))
}
