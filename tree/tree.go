package tree

import "fmt"

type Node struct {
	parent *Node
	value  string
	child  *Node
}

type Tree struct {
	Root       *Node
	MaxDepth   int
	MaxBreadth int
}

func New(c Config) *Tree {

	if err := c.Validate(); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	return &Tree{
		MaxDepth:   c.MaxDepth,
		MaxBreadth: c.MaxBreadth,
	}
}
