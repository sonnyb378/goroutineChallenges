package easy

import (
	"sync"
)

func Factorial(input int) int {

	// Challenge 2: Concurrent Factorial Calculator
	// Write a Go program that calculates the factorial of a given number using a
	// separate goroutine for the calculation.

	// Input:
	// n := 5

	// Desired Output:
	// Factorial of 5 is 120

	wg := sync.WaitGroup{}

	resultChan := make(chan int)

	wg.Add(1)
	go ComputeFactorial(input, &wg, resultChan)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	return <-resultChan

}

func ComputeFactorial(n int, wg *sync.WaitGroup, resultChan chan int) {
	fSum := 1
	defer wg.Done()
	if n <= 0 {
		resultChan <- 1
	}

	for i := n; i >= 1; i-- {
		fSum *= i
	}
	resultChan <- fSum
}
