package tree

import (
	"errors"
)

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
func Build(input []Record) (*Node, error) {
	if len(input) == 0 {
		return nil, nil
	}

	nodes := make([]Node, len(input))
	var root *Node = nil

	for i := 0; i < len(input); i++ {
		nodes[i] = Node{ID: i}
	}

	IDsSeen := map[int]bool{}
	for _, record := range input {
		ID := record.ID
		parent := record.Parent
		if IDsSeen[ID] {
			return nil, errors.New("Duplicate ID")
		}
		if ID < 0 || ID >= len(input) || parent < 0 || parent >= len(input) {
			return nil, errors.New("Bad ID")
		}
		if parent > ID {
			return nil, errors.New("Bad parent")
		}
		IDsSeen[ID] = true
		if ID == parent {
			if root != nil {
				return nil, errors.New("duplicate root")
			}
			root = &nodes[ID]
		} else {
			nodes[parent].Children = append(nodes[parent].Children, &(nodes[ID]))
		}
	}
	if root == nil {
		return nil, errors.New("no root")
	}

	return root, nil
}
