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

// returns a shallow copy of the row
func (this *Matrix) Row(index int) []float64 {
    return this.rows[index]
}

func (this *Matrix) SwapRows(i, j int) {
    temp := this.rows[i]
    this.rows[i] = this.rows[j]
    this.rows[j] = temp 
}

// Here I assume the row is of the correct size, or else woe be on ye
func (this *Matrix) PutRow(index int, newRow []float64) {
    if len(newRow) != this.W() {
        panic("Invalid row length")
    }
    copy(this.rows[index], newRow)
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
/* runs in N^3 time (for square matrices of size N x N), a.rows * a.cols * b.cols time for rectangular matrices */
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

func (this *Matrix) Submatrix(minCol, maxCol, minRow, maxRow int) *Matrix {
    k := make([][]float64, maxRow - minRow)  
    for i := range k {
        k[i] = this.Row(minRow + i)[minCol:maxCol]
    }
    m, _ := New(k)
    return m
}

// number of columns in the matrix
func (this *Matrix) W() int {
    return this.width
}

// number of rows in the matrix
func (this *Matrix) H() int {
    return this.height
}

func Identity(n int) *Matrix {
    m := make([][]float64, n)  
    for i := range m {
        m[i] = make([]float64, n)
        m[i][i] = 1.0
    }
    matrixM, _ := New(m)
    return matrixM 
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
