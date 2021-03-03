package raindrops

import (
	"fmt"
)

// Convert makes rain sounds
func Convert(n int) string {
	rainNoise := ""
	if n%3 == 0 {
		rainNoise += "Pling"
	}
	if n%5 == 0 {
		rainNoise += "Plang"
	}
	if n%7 == 0 {
		rainNoise += "Plong"
	}
	if rainNoise == "" {
		rainNoise = fmt.Sprint(n)
	}
	return rainNoise
}
