package tree

// Record stores the ids of the tree nodes
type Record struct {
	ID, Parent int
}

// Node is a node in the tree
type Node struct {
	ID       int
	Children []*Node
}

// Build builds a tree
func Build(input []Record) (Node, error) {
	return Node{}, nil
}
