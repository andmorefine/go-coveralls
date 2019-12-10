package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	result := Sample()
	assert.Equal(t, "name Grass", result)
}
