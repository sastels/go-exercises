package hamming

import (
	"errors"
)

// Distance computes Hamming distance
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("unequal lengths")
	}

	distance := 0
	for i, x := range a {
		if byte(x) != b[i] {
			distance++
		}

	}
	return distance, nil
}
