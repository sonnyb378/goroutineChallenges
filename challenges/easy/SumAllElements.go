package easy

import (
	"fmt"
	"sync"
	"time"
)

func SumAllElements(input []int) (string, error) {

	// Challenge 1: Simple Parallel Summation
	// Write a Go program that calculates the sum of all elements in a large array
	// concurrently using goroutines.

	// Input:
	// data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Desired Output:
	// Sum: 55

	wg := sync.WaitGroup{}

	numberOfGoroutines := 3
	chunkSize := len(input) / numberOfGoroutines
	var allSum int
	partialSum := make(chan int)

	timeStart := time.Now()

	for i := 0; i < numberOfGoroutines; i++ {

		start := i * chunkSize
		end := (i + 1) * chunkSize

		// For the last chunk, include any remaining elements
		if i == numberOfGoroutines-1 {
			end = len(input)
		}

		wg.Add(1)
		go SumAll(&wg, partialSum, input[start:end])
	}

	go func() {
		wg.Wait()
		close(partialSum)
	}()

	// Receive partialSum channel value
	for v := range partialSum {
		allSum += v
		time.Sleep(time.Second * 1)
	}

	elapsedTime := time.Since(timeStart)
	fmt.Println("Elapsed Time: ", elapsedTime.Round(time.Millisecond))
	return fmt.Sprintf("Sum: %v\n", allSum), nil
}

func SumAll(wg *sync.WaitGroup, ps chan int, input []int) {

	defer wg.Done()

	sum := 0
	for _, v := range input {
		sum += v
	}
	ps <- sum
	fmt.Println("Partial Sum Computation: ", sum)
}
