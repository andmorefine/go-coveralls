package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleAnimal(t *testing.T) {
	result := SampleAnimal()
	assert.Equal(t, "favorite Banana", result)
}
