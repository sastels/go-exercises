package grains

import "errors"

// Square returns the number of grains on square n
func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("bad value")
	}
	return 1 << (n - 1), nil
}

// Total returns the number of grains on the board
func Total() uint64 {
	sum := uint64(0)
	for i := 1; i < 65; i++ {
		n, _ := Square(i)
		sum += n
	}
	return sum
}
