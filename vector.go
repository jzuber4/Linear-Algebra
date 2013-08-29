package matrixmath

import "math"

func DotProduct(a, b []float64) float64 {
	if len(a) != len(b) {
		return math.NaN()
	}
	var v float64
	for i := range a {
		v += a[i] * b[i]
	}
	return v
}
