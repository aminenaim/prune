package tree

import "errors"

type Config struct {
	MaxDepth   int
	MaxBreadth int
}

func (c Config) Validate() error {

	if c.MaxDepth < 4 {
		return errors.New("config error: max depth should be greater or equal to 4")
	}

	if c.MaxBreadth < 4 {
		return errors.New("config error: max breadth should be greater or equal to 4")
	}
	return nil
}
