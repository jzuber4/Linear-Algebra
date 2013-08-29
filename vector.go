// Contains basic operations for float64 slices such as dot product, addition, and scaling. All operations return a new []float64 and do not alter their arguments.
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

func AddVectors(a, b []float64) []float64 {
	if len(a) != len(b) {
		return nil
	}
	newV := make([]float64, len(a))
	for i := range a {
		newV[i] = a[i] + b[i]
	}
	return newV
}

func SubtractVectors(a, b []float64) []float64 {
	if len(a) != len(b) {
		return nil
	}
	newV := make([]float64, len(a))
	for i := range a {
		newV[i] = a[i] - b[i]
	}
	return newV
}

func ScaleVector(oldV []float64, s float64) []float64 {
	newV := make([]float64, len(oldV))
	copy(newV, oldV)
	for i, val := range newV {
		newV[i] = val * s
	}
	return newV
}

func DivideVector(oldV []float64, s float64) []float64 {
	newV := make([]float64, len(oldV))
	copy(newV, oldV)
	for i, val := range newV {
		newV[i] = val / s
	}
	return newV
}
