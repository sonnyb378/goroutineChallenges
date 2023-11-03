package hard

import (
	"fmt"
	"sync"
)

func ProducerConsumer(items int) {

	// Challenge 9: Implement the Producer-Consumer problem using Mutex.
	// Description:
	// In this challenge, you are tasked with implementing a producer-consumer pattern in Go. The producer-consumer pattern involves two types of goroutines:

	// Producers: These goroutines generate items and add them to a shared buffer.
	// Consumers: These goroutines retrieve items from the shared buffer and process them.

	// Output:
	// The expected output should demonstrate that producers produce items and consumers consume them. The relative order of production and
	// consumption may vary due to the concurrent nature of the program. However, it is important to verify that all items are produced and consumed,
	// and that proper synchronization is in place to prevent data races.

	// Here are sample output snippets that represent the correct behavior:

	// Produced: Item 0
	// Consumed: Item 0
	// Produced: Item 1
	// Consumed: Item 1
	// Produced: Item 2
	// Consumed: Item 2
	// Produced: Item 3
	// Consumed: Item 3
	// Produced: Item 4
	// Consumed: Item 4
	// ...

	wg := sync.WaitGroup{}
	rw := sync.RWMutex{}
	itemsCh := make(chan int, items)

	for i := 0; i < items; i++ {
		wg.Add(1)
		go func(item int) {
			defer wg.Done()
			SafeReadWrite(&rw, func() {
				fmt.Printf("Produced: Item %v\n", item)
				itemsCh <- item
			})
		}(i)
	}

	go func() {
		wg.Wait()
		close(itemsCh)
	}()

	for item := range itemsCh {
		fmt.Printf("Consumed: Item %v\n", item)
	}

}

func SafeReadWrite(rw *sync.RWMutex, operation func()) {
	rw.RLock()
	defer rw.RUnlock()
	operation()
}
