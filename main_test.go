package main

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name     string
	Language int
	errorMtg int
}

func TestMain(m *testing.M) {
	func() {
		fmt.Println("Prepare test")
	}()
	ret := m.Run()
	func() {
		fmt.Println("Teardown test")
	}()

	os.Exit(ret)
}

func TestExampleSuccess(t *testing.T) {
	cases := []struct {
		st       string
		expected int
		errorMtg error
	}{
		{"hoge", 1, nil},
		{"moge", 0, errors.New("code must be hoge")},
	}

	for _, testCase := range cases {
		result, ok := example(testCase.st)
		assert.Equal(t, result, testCase.expected)
		assert.Equal(t, ok, testCase.errorMtg)
	}
}
