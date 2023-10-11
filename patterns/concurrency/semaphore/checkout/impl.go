package checkout

import (
	"fmt"
	"sync"
	"time"
)

var (
	semaphore = make(chan struct{}, 2) // Define the semaphore to 2
	wg        sync.WaitGroup
)

// Process simulates processing an e-commerce checkout for a given order ID.
func Process(orderID int) {
	semaphore <- struct{}{}        // Acquire a Lock
	defer func() { <-semaphore }() // Release a Lock

	fmt.Printf("Processing order %d\n", orderID)
	time.Sleep(1 * time.Second) // Simulate payment processing
	fmt.Printf("Order %d processed\n", orderID)
	wg.Done()
}

// ProcessOrders processes a slice of order IDs.
func ProcessOrders(orderIDs []int) {
	for _, orderID := range orderIDs {
		wg.Add(1)
		go Process(orderID)
	}

	wg.Wait()
	fmt.Println("All orders processed")
}
