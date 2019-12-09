package main

import (
	"errors"
	"fmt"
)

func main() {
	result, _ := example("mo")
	fmt.Println(result)
}

func example(code string) (int, error) {
	if code == "hoge" {
		return 1, nil
	}
	return 0, errors.New("code must be hoge")
}
