package intermediate

import (
	"fmt"
	"sync"
)

func BankTransactions(startingBalance float64, numberOfCustomers int, numberOfTransactions int) {
	// Challenge 8: Bank Account Transactions with Mutex
	// Title: Simulate bank account transactions with Mutex.

	// Description: Create a Go program that simulates bank account transactions by multiple concurrent goroutines.
	// Each goroutine represents a customer and can either deposit or withdraw a random amount of money from the account.
	// Use a Mutex to protect the shared bank account data. Ensure that the account balance is updated correctly and does not go negative.

	// Input: Number of customers, initial account balance, and the number of transactions each customer should perform.

	// Desired Output: The final balance of the bank account after all transactions.

	// Correct Output:
	// Final Balance: $5000.00

	// Wrong Output:
	// Final Balance: $-126.75 (due to data races)

	var initialBalance float64
	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	initialBalance = startingBalance

	amount := 20.00

	for i := 0; i < numberOfCustomers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < numberOfTransactions; j++ {
				if j%2 == 0 {
					SafeWrite(&m, initialBalance, func() {
						initialBalance = initialBalance + amount
					})
				} else {
					SafeWrite(&m, initialBalance, func() {
						initialBalance = initialBalance - amount
					})
				}
			}

		}(i)
	}

	wg.Wait()
	fmt.Printf("Final Balance: $%v\n", initialBalance)
}

func SafeWrite(m *sync.Mutex, balance float64, operation func()) {
	m.Lock()
	defer m.Unlock()
	operation()
}
