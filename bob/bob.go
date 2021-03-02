package bob

import (
	"strings"
	"unicode"
)

// Hey implements a teenager
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	hasLetters := false
	for _, r := range strings.ToUpper(remark) {
		if unicode.IsLetter(r) {
			hasLetters = true
		}
	}
	isYell := hasLetters && (strings.ToUpper(remark) == remark)
	isQuestion := strings.HasSuffix(remark, "?")
	isEmpty := remark == ""

	switch {
	case isEmpty:
		return "Fine. Be that way!"
	case isQuestion && isYell:
		return "Calm down, I know what I'm doing!"
	case isQuestion:
		return "Sure."
	case isYell:
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
