package main

import (
	"fmt"

	"github.com/sonnyb378/goroutineChallenges/challenges/easy"
)

func main() {

	// EASY

	// fmt.Println("Challenge 1: Simple Parallel Summation")
	// data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// result, err := easy.SumAllElements(data)
	// fmt.Println(">>> ", result, err)
	// fmt.Println("=========================================")
	// fmt.Println("")

	fmt.Println("Challenge 2: Concurrent Factorial Calculator")
	num := 5
	factorial := easy.Factorial(num)
	fmt.Printf("Factorial of %v is %v\n", num, factorial)
	fmt.Println("=========================================")
	fmt.Println("")
}
