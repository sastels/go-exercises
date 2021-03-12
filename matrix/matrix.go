package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix represents a matrix
type Matrix struct {
	data             []int
	numRows, numCols int
}

// New returns a new matrix
func New(input string) (*Matrix, error) {
	var matrix Matrix
	for _, line := range strings.Split(input, "\n") {
		matrix.numRows++
		line = strings.TrimSpace(line)
		lineSplit := strings.Split(line, " ")
		if matrix.numCols != 0 && matrix.numCols != len(lineSplit) {
			return nil, errors.New("inconsistant width")
		}
		matrix.numCols = len(lineSplit)
		for _, nString := range lineSplit {
			n, err := strconv.Atoi(nString)
			if err != nil {
				return nil, err
			}
			matrix.data = append(matrix.data, n)
		}
	}
	return &matrix, nil
}

// Set sets part of the matrix
func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= m.numRows || col < 0 || col >= m.numCols {
		return false
	}
	m.data[row*m.numCols+col] = val
	return true
}

// Rows returns the rows of the matrix
func (m Matrix) Rows() [][]int {
	rows := [][]int{}
	for y := 0; y < m.numRows; y++ {
		row := []int{}
		for x := 0; x < m.numCols; x++ {
			row = append(row, m.data[y*m.numCols+x])
		}
		rows = append(rows, row)
	}
	return rows
}

// Cols returns the columns of the matrix
func (m Matrix) Cols() [][]int {
	cols := [][]int{}
	for x := 0; x < m.numCols; x++ {
		col := []int{}
		for y := 0; y < m.numRows; y++ {
			col = append(col, m.data[y*m.numCols+x])
		}
		cols = append(cols, col)
	}
	return cols
}
