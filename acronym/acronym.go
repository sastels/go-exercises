package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate makes an acronym
func Abbreviate(s string) string {
	matches := regexp.MustCompile(`([-_])`)
	s = matches.ReplaceAllString(s, " ")

	matches = regexp.MustCompile(`(\s+)`)
	s = matches.ReplaceAllString(s, " ")

	acc := ""
	for _, value := range strings.Split(s, " ") {
		acc += strings.ToUpper(value[:1])
	}
	return acc
}
