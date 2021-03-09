package tree

import (
	"errors"
	"fmt"
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

	for _, record := range input {
		if record.ID < 0 || record.ID >= len(input) || record.Parent < 0 || record.Parent >= len(input) {
			return nil, errors.New("Bad ID")
		}

		fmt.Printf("id %d, parent %d\n", record.ID, record.Parent)
		if record.ID == record.Parent {
			root = &nodes[record.ID]
		} else {
			thisNode := &(nodes[record.ID])
			nodes[record.Parent].Children = append(nodes[record.Parent].Children, thisNode)
		}
	}

	if root == nil {
		return nil, errors.New("no root")
	}
	return root, nil
}
