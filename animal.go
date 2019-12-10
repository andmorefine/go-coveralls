package main

import (
	"fmt"
	"go-coveralls/animals"
)

// SampleAnimal name
func SampleAnimal() string {
	return fmt.Sprintf("favorite %s", animals.MonkeyFeed())
}
