package core

import (
	"fmt"

	"github.com/aminenaim/prune/tree"
)

func Run() {
	t := tree.New(tree.Config{
		MaxDepth:   1,
		MaxBreadth: 4,
	})

	fmt.Printf("MaxDepth: %d\n", t.MaxDepth)
	fmt.Printf("MaxBreadth: %d\n", t.MaxBreadth)
}
