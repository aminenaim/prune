package core

import (
	"fmt"

	"github.com/aminenaim/prune/tree"
)

// InitTree Initializes a tree with the given maxDepth and maxBreadth
func InitTree(maxDepth int, maxBreadth int) *tree.Tree {
	c := tree.Config{
		MaxDepth:   maxDepth,
		MaxBreadth: maxBreadth,
	}

	return tree.New(c)
}

// Run Runs the core package
func Run() {
	t := InitTree(4, 4)

	fmt.Printf("MaxDepth: %d\n", t.MaxDepth)
	fmt.Printf("MaxBreadth: %d\n", t.MaxBreadth)
}
