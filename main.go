package main

import (
	"fmt"

	"github.com/sonnyb378/goroutineChallenges/challenges/easy"
	// "github.com/sonnyb378/goroutineChallenges/challenges/intermediate"
	// "github.com/sonnyb378/goroutineChallenges/challenges/hard"
)

func main() {

	// EASY

	// fmt.Println("Challenge 1: Simple Parallel Summation")
	// data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// result, err := easy.SumAllElements(data)
	// fmt.Println(">>> ", result, err)
	// fmt.Println("=========================================")
	// fmt.Println("")

	// fmt.Println("Challenge 2: Concurrent Factorial Calculator")
	// num := 5
	// factorial := easy.Factorial(num)
	// fmt.Printf("Factorial of %v is %v\n", num, factorial)
	// fmt.Println("=========================================")
	// fmt.Println("")

	fmt.Println("Challenge: Concurrent Counter")
	easy.Counter(10, 1000)

	// INTERMEDIATE
	// fmt.Println("Challenge 3: Concurrent Prime Number Finder")
	// input := []int{
	// 	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	// 	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	// 	21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	// 	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	// 	41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
	// 	51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
	// 	61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
	// 	71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
	// 	81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
	// 	91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
	// }
	// result, err := intermediate.FindPrimes(input, 5)
	// fmt.Println("Result: ", result, err)
	// fmt.Println("=========================================")
	// fmt.Println("")

	// fmt.Println("Challenge 4: Concurrent Word Count")
	// path := "challenges/intermediate/documents"
	// result, err := intermediate.WordCountInFile(path)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }
	// for _, v := range result {
	// 	fmt.Println(v)
	// }
	// fmt.Println("=========================================")
	// fmt.Println("")

	// HARD
	// fmt.Println("Challenge 5: Parallel Matrix Multiplication")
	// result := hard.MatrixMultiplication(3, 3)
	// fmt.Println("Final Product Matrix: ", result)

	// fmt.Println("Challenge 6: Concurrent Sudoku Solver")
	// hard.Sudoku()

}
