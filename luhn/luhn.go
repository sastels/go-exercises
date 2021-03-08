package luhn

import (
	"strings"
	"unicode"
)

// Valid checks if s is a Luan number
func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	if len(s) <= 1 {
		return false
	}
	sum := 0
	for i, r := range s {
		if unicode.IsDigit(r) {
			n := int(r - '0')
			if (len(s)-i)%2 == 0 {
				sum += 2 * n
				if 2*n > 9 {
					sum -= 9
				}
			} else {
				sum += n
			}
		} else {
			return false
		}
	}
	return sum%10 == 0
}
