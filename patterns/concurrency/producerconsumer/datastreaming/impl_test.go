package datastreaming

import (
	"reflect"
	"sync"
	"testing"
)

func TestProducerConsumer(t *testing.T) {
	tests := []struct {
		name           string
		itemsToProduce []Item
		expected       []Item
	}{
		{
			name:           "Test with 5 items",
			itemsToProduce: []Item{{1}, {2}, {3}, {4}, {5}},
			expected:       []Item{{2}, {4}, {6}, {8}, {10}},
		},
		{
			name:           "Test with 3 items",
			itemsToProduce: []Item{{6}, {7}, {8}},
			expected:       []Item{{12}, {14}, {16}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan Item, 10)
			var wg sync.WaitGroup
			processedItems := make([]Item, 0)

			wg.Add(1)
			go func() {
				Producer(ch, &wg, tt.itemsToProduce)
				close(ch) // Close the channel after the producer is done sending items.
			}()

			wg.Add(1)
			go Consumer(ch, &wg, &processedItems)

			wg.Wait() // Wait for both the producer and consumer to finish.

			if !reflect.DeepEqual(processedItems, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, processedItems)
			}
		})
	}
}
