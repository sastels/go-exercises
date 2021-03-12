package sublist

// Relation is relationship between lists "equal", "unequal", "sublist", "superlist"
type Relation string

// a contains b
func contains(a []int, b []int) bool {
	if len(b) == 0 {
		return true
	}
	if len(b) > len(a) {
		return false
	}
	for aStart := 0; aStart <= len(a)-len(b); aStart++ {
		aContainsB := true
		for index := 0; index < len(b); index++ {
			if a[aStart+index] != b[index] {
				aContainsB = false
				break
			}
		}
		if aContainsB {
			return true
		}
	}
	return false
}

// Sublist returns the sublist relationship between the lists
func Sublist(a []int, b []int) Relation {
	switch {
	case len(a) == len(b) && contains(a, b):
		return "equal"
	case contains(a, b):
		return "superlist"
	case contains(b, a):
		return "sublist"
	}
	return "unequal"
}
