package tree

import "errors"

// Config is the configuration for the program.
//
// - MaxDepth is the maximum depth of the tree
//
// - MaxBreadth is the maximum breadth of the tree
type Config struct {
	MaxDepth   int
	MaxBreadth int
}

// Validate checks if the configuration is valid
//
// - Config is the configuration for the program.
func (c Config) Validate() error {

	if c.MaxDepth < 4 {
		return errors.New("config error: max depth should be greater or equal to 4")
	}

	if c.MaxBreadth < 4 {
		return errors.New("config error: max breadth should be greater or equal to 4")
	}
	return nil
}
