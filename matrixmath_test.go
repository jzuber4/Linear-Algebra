// Contains definition of matrix type as well as basic operations for matrices including addition, subtraction, and multiplication
package matrixmath

import (
	"math/rand"
	"testing"
	"time"
)

// create a random 2d slice of length h containing slices of length w
// slices contain floats [-max/2, max/2)
func random2dSlice(w, h int, max float64) [][]float64 {
	a := make([][]float64, h)
	for i := range a {
		a[i] = make([]float64, w)
		for j := range a[i] {
			a[i][j] = rand.Float64()*max - (max / 2.0)
		}
	}
	return a
}

// Test matrix addition
func TestAdd(t *testing.T) {
	// seed some randomness
	rand.Seed(time.Now().UTC().UnixNano())
	// matrices with length/widths of 1 to 100
	w := rand.Intn(99) + 1
	h := rand.Intn(99) + 1
	max := 100.0
	a := random2dSlice(w, h, max)
	b := random2dSlice(w, h, max)
	// perform our own addition
	c := make([][]float64, h)
	for i := range a {
		c[i] = make([]float64, w)
		for j := range a[i] {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	// error check each New()
	matrixA, e := New(a)
	if e != nil {
		t.Errorf("%v", e.Error())
	}
	matrixB, e := New(b)
	if e != nil {
		t.Errorf("%v", e.Error())
	}
	matrixC1, e := New(c)
	if e != nil {
		t.Errorf("%v", e.Error())
	}
	// Our packages matrix add
	matrixC2, e := Add(matrixA, matrixB)
	if e != nil {
		t.Errorf("%v", e.Error())
	}
	// Test if the addition in this function equals the package addition
	if !Equal(matrixC1, matrixC2) {
		t.Errorf("Add operation produced incorrect result")
	}

}

// Test matrix subtraction
func TestSubtract(t *testing.T) {
	// seed some randomness
	rand.Seed(time.Now().UTC().UnixNano())
	// matrices
	w := rand.Intn(99) + 1
	h := rand.Intn(99) + 1
	max := 100.0
	a := random2dSlice(w, h, max)
	b := random2dSlice(w, h, max)
	c := make([][]float64, h)
	// perform our own subtraction
	for i := range a {
		c[i] = make([]float64, w)
		for j := range a[i] {
			c[i][j] = a[i][j] - b[i][j]
		}
	}
	matrixA, e := New(a)
	if e != nil {
		t.Errorf("%v", e.Error())
	}
	matrixB, e := New(b)
	if e != nil {
		t.Errorf("%v", e.Error())
	}
	matrixC1, e := New(c)
	if e != nil {
		t.Errorf("%v", e.Error())
	}
	// subtraction to be tested
	matrixC2, e := Subtract(matrixA, matrixB)
	if e != nil {
		t.Errorf("%v", e.Error())
	}
	// compare subtractions
	if !Equal(matrixC1, matrixC2) {
		t.Errorf("Subtract operation produced incorrect result")
	}
}

func TestMultiply(t *testing.T) {
	// seed some randomness
	rand.Seed(time.Now().UTC().UnixNano())
	// test with matrices of width/length 1 to 100
	w := rand.Intn(99) + 1
	h := rand.Intn(99) + 1
	w2 := rand.Intn(99) + 1
	max := 100.0
	a := random2dSlice(w, h, max)
	b := random2dSlice(w2, w, max)
	c := make([][]float64, h, h)
	matrixA, e := New(a)
	if e != nil {
		t.Errorf("%v", e.Error())
	}
	matrixB, e := New(b)
	if e != nil {
		t.Errorf("%v", e.Error())
	}
	matrixC2, e := Multiply(matrixA, matrixB)
	if e != nil {
		t.Errorf("%v", e.Error())
	}
	// A * B = dotproduct (row i of A, column j of B) for each cell i, j
	for i := range c {
		c[i] = make([]float64, w2)
		for j := range c[i] {
			c[i][j] = DotProduct(a[i], matrixB.Col(j))
		}
	}
	matrixC1, e := New(c)
	if e != nil {
		t.Errorf("%v", e.Error())
	}
	if !Equal(matrixC1, matrixC2) {
		t.Errorf("Multiply operation produced incorrect result")
	}
}

func TestEqual(t *testing.T) {
	// seed some randomness
	rand.Seed(time.Now().UTC().UnixNano())
	w0 := rand.Intn(99) + 1
	h0 := rand.Intn(99) + 1
	w1 := w0 + rand.Intn(5) + 1
	h1 := h0 + rand.Intn(5) + 1
	// base
	a := random2dSlice(w0, h0, 100.0)
	// same size, different contents
	b := random2dSlice(w0, h0, 100.0)
	b[h0-1][w0-1] = a[h0-1][w0-1] + 1
	// same size, same contents
	c := make([][]float64, h0)
	for i := range c {
		c[i] = make([]float64, w0)
		copy(c[i], a[i])
	}
	// larger height, otherwise same contents
	d := make([][]float64, h1)
	for i := range d {
		d[i] = make([]float64, w0)
		if i < h0 {
			copy(d[i], a[i])
		}
	}
	// larger width, otherwise same contents
	e := make([][]float64, h0)
	for i := range c {
		e[i] = make([]float64, w1)
		copy(e[i], a[i])
	}
	matrixA, _ := New(a)
	matrixB, _ := New(b)
	matrixC, _ := New(c)
	matrixD, _ := New(d)
	matrixE, _ := New(e)
	if !Equal(matrixA, matrixA) {
		t.Errorf("Equal returned false for two pointers to the same matrix")
	}
	if Equal(matrixA, matrixB) || Equal(matrixB, matrixA) {
		t.Errorf("Equal returned true for matrices with different contents and the same size")
	}
	if !Equal(matrixA, matrixC) || !Equal(matrixC, matrixA) {
		t.Errorf("Equal returned false for two matrices with the same contents.", matrixA, matrixC)
	}
	if Equal(matrixA, matrixD) || Equal(matrixD, matrixA) {
		t.Errorf("Equal returned true for matrices of different height")
	}
	if Equal(matrixA, matrixE) || Equal(matrixE, matrixA) {
		t.Errorf("Equal returned true for matrices of different width")
	}

}
