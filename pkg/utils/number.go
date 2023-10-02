package utils

import "math/rand"

func GetMinFloat64(a, b float64) float64 {
	if a < b {
		return a
	}

	return b
}

func RandomForMinFloat64(a, b float64) float64 {
	return GetMinFloat64(a, b) * rand.Float64()
}
