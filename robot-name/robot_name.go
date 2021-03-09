package robotname

import (
	"errors"
	"math/rand"
	"time"
)

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var digits = []rune("0123456789")

var namesSeen = map[string]bool{}

// Robot is a robot struct
type Robot struct {
	name string
}

// Name returns the name of the robot
func (r *Robot) Name() (string, error) {
	rand.Seed(time.Now().UnixNano())

	if r.name == "" {
		if len(namesSeen) == 26*26*10*10*10 {
			return "", errors.New("Namespace exhausted")
		}
		for {
			nameRunes := make([]rune, 5)
			nameRunes[0] = letters[rand.Intn(len(letters))]
			nameRunes[1] = letters[rand.Intn(len(letters))]
			nameRunes[2] = digits[rand.Intn(len(digits))]
			nameRunes[3] = digits[rand.Intn(len(digits))]
			nameRunes[4] = digits[rand.Intn(len(digits))]
			r.name = string(nameRunes)
			if namesSeen[r.name] == false {
				break
			}
		}
		namesSeen[r.name] = true
	}
	return r.name, nil
}

// Reset clears the robot's name
func (r *Robot) Reset() {
	r.name = ""
}
