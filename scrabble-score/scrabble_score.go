package scrabble

import "strings"

// Score computes the scrabble score of a word
func Score(s string) int {
	s = strings.ToUpper(s)
	score := 0
	var letterScore = map[string]int{
		"A, E, I, O, U, L, N, R, S, T": 1,
		"D, G":                         2,
		"B, C, M, P":                   3,
		"F, H, V, W, Y":                4,
		"K":                            5,
		"J, X":                         8,
		"Q, Z":                         10,
	}
	for _, c := range s {
		for key := range letterScore {
			if strings.ContainsRune(key, c) {
				score += letterScore[key]
			}
		}
	}
	return score
}
