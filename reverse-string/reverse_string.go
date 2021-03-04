package reverse

// Reverse reverses a string
func Reverse(s string) string {
	out := make([]rune, 0, len(s))
	for _, c := range s {
		out = append(out, 'x')
		copy(out[1:], out[0:])
		out[0] = c
	}
	return string(out)
}
