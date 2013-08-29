package matrixmath

import () 
/*
// Returns the inverse of the matrix m
// nil if m is non-invertible
func Invert(m *Matrix) *Matrix {
    return nil
}
*/

// Returns the reduced row echelon form of the matrix m
func RREF(oldM *Matrix) *Matrix {
    // copy the matrix
    m, _ := New(oldM.Contents())
    // index of reducing row in matrix 
    lastRow := -1
    for i := 0; i < m.W(); i++ {
        coefficient, row := 0.0, -1
        // Find row to use for reducing (has leading coefficient in column)
        for j := 0; j < m.H(); j++ {
            c, index := LeadingCoefficient(m.Row(j)) 
            if index == i {
                coefficient = c
                row = j
                break
            }
        }
        if coefficient == 0 {
            // none found
            continue
        }
        lastRow++
        // move the row into position
        m.SwapRows(lastRow, row)
        // divide the row by the leading coefficient
        m.PutRow(lastRow, DivideVector(m.Row(lastRow), coefficient))
        // subtract this row from every other row 
        for j := 0; j < m.H(); j++ {
            if j == lastRow {
                continue 
            }
            if m.Row(j)[i] != 0 {
                m.PutRow(j, SubtractVectors(m.Row(j), ScaleVector(m.Row(lastRow), m.Row(j)[i])))
            }
        }
    }
    return m
}

/*
// Returns vector(s) solving the equations defined by the matrix m
// nil if there is no solution
func SolveAsLinearEquations(m *Matrix) [][]float64 {

    return nil
}

// Returns a matrix that spans the kernel of matrix m
func Kernel(m *Matrix) *Matrix {

    return nil
}

// Returns a matrix that spans the image of matrix m
func Basis(m *Matrix) *Matrix {

    return nil
}

*/
