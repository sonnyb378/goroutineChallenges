package hard

func ProducerConsumer() {
	// Challenge: Producer-Consumer Problem with Mutex
	// Input: Number of items to produce and consume: 10

	// Desired Correct Output:
	// Produced: Item 1
	// Consumed: Item 1
	// Produced: Item 2
	// Consumed: Item 2
	// Produced: Item 3
	// Consumed: Item 3
	// ...

	// Wrong Output (due to data races or incorrect synchronization):
	// Produced: Item 1
	// Consumed: Item 1
	// Produced: Item 2
	// Produced: Item 3
	// Consumed: Item 2
	// Consumed: Item 3
	// ...

	panic("")
}
