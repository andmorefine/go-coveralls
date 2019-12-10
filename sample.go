package main

import (
	"fmt"
	"go-coveralls/animals"
)

// Sample name
func Sample() string {
	return fmt.Sprintf("name %s", animals.ElephantFeed())
}
