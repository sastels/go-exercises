package sublist

// Relation is relationship between lists
type Relation string

// Sublist returns the sublist relationship between the lists
func Sublist(a []int, b []int) Relation {
	return "none"
}
