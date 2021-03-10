package twelve

import "fmt"

var numbers = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

var gifts = map[int]string{
	1:  "a Partridge in a Pear Tree",
	2:  "two Turtle Doves",
	3:  "three French Hens",
	4:  "four Calling Birds",
	5:  "five Gold Rings",
	6:  "six Geese-a-Laying",
	7:  "seven Swans-a-Swimming",
	8:  "eight Maids-a-Milking",
	9:  "nine Ladies Dancing",
	10: "ten Lords-a-Leaping",
	11: "eleven Pipers Piping",
	12: "twelve Drummers Drumming",
}

// Verse is one verse of the song
func Verse(n int) string {
	ret := fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s", numbers[n], gifts[n])
	for i := n - 1; i > 1; i-- {
		ret = fmt.Sprintf("%s, %s", ret, gifts[i])
	}
	if n > 1 {
		ret = fmt.Sprintf("%s, and %s", ret, gifts[1])
	}
	ret += "."
	return ret
}

// Song returns the whole song
func Song() string {
	ret := ""
	for i := 1; i <= 12; i++ {
		ret += Verse(i)
		if i != 12 {
			ret += "\n"
		}
	}
	return ret
}
