package intermediate

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/sonnyb378/goroutineChallenges/pkg/customerror"
)

func FindPrimes(input []int, numberOfGoroutines int) ([]int, error) {

	// Challenge 1: Concurrent Prime Number Finder
	// Write a Go program that finds and prints all prime numbers in a given range concurrently.
	// The program should accept a range of numbers, divide the range into segments, and use
	// goroutines to check for prime numbers within each segment.
	// Print the prime numbers as they are discovered.

	// Input:
	// A range of numbers, e.g., from 1 to 100.

	// Desired Output:
	// Prime numbers in the range 1-100:
	// 2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97

	wg := sync.WaitGroup{}
	// rm := sync.RWMutex{}
	primeNumbers := []int{}

	segment := numberOfGoroutines
	if numberOfGoroutines > len(input) {
		return []int{}, &customerror.CustomError{
			Message: "Invalid chunk size!",
		}
	}

	chunkSize := len(input) / segment
	ch := make(chan int)

	for i := 0; i < segment; i++ {

		start := i * chunkSize
		end := start + chunkSize

		if i == segment-1 {
			end = len(input)
		}
		wg.Add(1)

		go func() {
			defer wg.Done()
			timeStart := time.Now()
			for _, v := range input[start:end] {
				if IsPrime(v) {
					ch <- v
				}
			}
			elapsedTime := time.Since(timeStart)
			fmt.Printf("Execution time: %v (in Microseconds)\n", elapsedTime.Microseconds())

		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for c := range ch {
		primeNumbers = append(primeNumbers, c)
	}
	sort.Ints(primeNumbers)
	return primeNumbers, nil
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
