package datastreaming

import "sync"

type Item struct {
	Data int
}

// Producer sends items to a channel.
func Producer(ch chan<- Item, wg *sync.WaitGroup, items []Item) {
	defer wg.Done()
	for _, item := range items {
		ch <- item
	}
}

// Consumer reads from a channel, processes the items (multiplies the data by 2), and appends them to a slice.
func Consumer(ch <-chan Item, wg *sync.WaitGroup, processedItems *[]Item) {
	defer wg.Done()
	for item := range ch {
		// Simulate processing
		item.Data *= 2
		*processedItems = append(*processedItems, item)
	}
}
