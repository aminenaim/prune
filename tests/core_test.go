package tests

import (
	"github.com/aminenaim/prune/core"
	"testing"
)

func TestInitTree(t *testing.T) {
	tree := core.InitTree(4, 4)
	if tree.MaxDepth != 4 {
		t.Errorf("Expected MaxDepth to be 4, got %d", tree.MaxDepth)
	}

	if tree.MaxBreadth != 4 {
		t.Errorf("Expected MaxBreadth to be 4, got %d", tree.MaxBreadth)
	}
}
