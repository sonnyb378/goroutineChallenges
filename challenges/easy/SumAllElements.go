package easy

import (
	"fmt"
)

func SumAllElements(input []int) (string, error) {

	// Challenge 1: Simple Parallel Summation
	// Write a Go program that calculates the sum of all elements in a large array
	// concurrently using goroutines.

	// Input:
	// data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Desired Output:
	// Sum: 55

	var allSum int
	return fmt.Sprintf("Sum: %v\n", allSum), nil
}
