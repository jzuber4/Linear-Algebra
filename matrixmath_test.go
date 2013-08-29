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
		t.Errorf("Add(%v, %v) = %v, want %v", matrixA, matrixB, matrixC2, matrixC1)
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
		t.Errorf("Subtract(%v, %v) = %v, want %v", matrixA, matrixB, matrixC2, matrixC1)
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
	matrixC1, e := Multiply(matrixA, matrixB)
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
	matrixC2, e := New(c)
	if e != nil {
		t.Errorf("%v", e.Error())
	}
	if !Equal(matrixC1, matrixC2) {
		t.Errorf("Add(%v, %v) = %v, want %v", matrixA, matrixB, matrixC2, matrixC1)
	}
}
