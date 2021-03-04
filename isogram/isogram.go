package isogram

import (
	"regexp"
	"strings"
)

// IsIsogram checks if a string has a repeated letter
func IsIsogram(s string) bool {
	matches := regexp.MustCompile(`([- ])`)
	s = matches.ReplaceAllString(s, "")
	s = strings.ToUpper(s)

	for i, c := range s {
		if strings.ContainsRune(s[:i], c) {
			return false
		}
	}
	return true
}
