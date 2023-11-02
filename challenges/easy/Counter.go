package easy

import (
	"fmt"
	"sync"
)

func Counter(numberOfGoroutines int, processTimes int) {

	// Challenge: Implement a concurrent counter using a Mutex.
	// Description: Write a Go program that creates multiple goroutines to increment a shared counter concurrently using a Mutex.
	// The program should increment the counter a specified number of times and ensure that the counter is protected from data races using a Mutex.
	// At the end of the program, print the final value of the counter.

	// Input: Number of goroutines and the number of times each goroutine should increment the counter.

	// Desired Output: The final value of the counter after all goroutines have finished.

	// Correct Output:
	// Counter Value: 10000

	// Wrong Output:
	// Counter Value: 9900 (due to data races)

	var counter int
	wg := &sync.WaitGroup{}
	m := sync.Mutex{}

	for i := 0; i < numberOfGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < processTimes; j++ {
				m.Lock()
				counter++
				m.Unlock()
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Counter: ", counter)
}
