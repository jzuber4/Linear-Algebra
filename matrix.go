package matrixmath

import ()

type Matrix struct {
	rows   [][]float64
	width  int
	height int
}

type matrixError struct {
	s string
}

func (e *matrixError) Error() string {
	return e.s
}

func New(input [][]float64) (*Matrix, *matrixError) {
	var m Matrix
	m.rows = make([][]float64, len(input), len(input))
	for i := range input {
		if i == 0 {
			continue
		}
		if len(input[i-1]) != len(input[i]) {
			return nil, &matrixError{"All vectors must be of same length"}
		}
	}
	for i := range input {
		m.rows[i] = make([]float64, len(input[i]), len(input[i]))
		copy(m.rows[i], input[i])
	}
	m.height = len(m.rows)
	m.width = len(m.rows[0])
	return &m, nil
}

func (this *Matrix) Col(index int) []float64 {
	col := make([]float64, this.height, this.height)
	for i := range this.rows {
		col[i] = this.rows[i][index]
	}
	return col
}

/* For each cell of matrices a and b, evaluates a function f and
stores the result in the same cell of a new matrix. Returns the new
matrix m. */
func Map2(a, b *Matrix, f func(a, b float64) float64) (*Matrix, *matrixError) {
	if a.width != b.width || a.height != b.height {
		return nil, &matrixError{"Matrices are of different sizes."}
	}
	m := make([][]float64, a.height, a.height)
	for i := range a.rows {
		m[i] = make([]float64, a.width, a.width)
		for j := range a.rows[i] {
			m[i][j] = f(a.rows[i][j], b.rows[i][j])
		}
	}
	return New(m)
}

func Add(a, b *Matrix) (*Matrix, *matrixError) {
	f := func(a, b float64) float64 {
		return a + b
	}
	return Map2(a, b, f)
}

func Subtract(a, b *Matrix) (*Matrix, *matrixError) {
	f := func(a, b float64) float64 {
		return a - b
	}
	return Map2(a, b, f)
}

func (this *Matrix) Contents() [][]float64 {
	contents := make([][]float64, this.height)
	for i := range contents {
		contents[i] = make([]float64, this.width)
		copy(contents[i], this.rows[i])
	}
	return contents
}

/* cache friendly multiplication of two matrices */
/* runs in N^3 time (for square matrices of size N x N), A.rows * a.cols * b.cols time for rectangular matrices */
func Multiply(a, b *Matrix) (*Matrix, *matrixError) {
	if a.width != b.height {
		return nil, &matrixError{"Matrix a must have the same number of columns as matrix b has rows."}
	}
	m := make([][]float64, a.height, a.height)
	for i := range m {
		m[i] = make([]float64, b.width, b.width)
	}
	for k := 0; k < a.width; k++ {
		for i := 0; i < a.height; i++ {
			for j := 0; j < b.width; j++ {
				m[i][j] = m[i][j] + a.rows[i][k]*b.rows[k][j]
			}
		}
	}
	return New(m)
}

func Equal(a, b *Matrix) bool {
    if a == b {
        return true
    }
	if a.height != b.height || a.width != b.width {
		return false
	}
	for i := range a.rows {
		for j := range a.rows[i] {
			if a.rows[i][j] != b.rows[i][j] {
				return false
			}
		}
	}
	return true
}
