package matrix

// Matrix represents a matrix
type Matrix struct {
	data             []int
	numRows, numCols int
}

// New returns a new matrix
func New(input string) (*Matrix, error) {
	var matrix Matrix
	return &matrix, nil
}

// Set sets part of the matrix
func (m *Matrix) Set(x, y, val int) bool {
	return false
}

// Rows returns the rows of the matrix
func (m Matrix) Rows() [][]int {
	return [][]int{}
}

// Cols returns the columns of the matrix
func (m Matrix) Cols() [][]int {
	return [][]int{}
}
