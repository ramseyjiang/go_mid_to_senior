package datastreaming

import (
	"reflect"
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
			// Initialize and Start
			stream := NewStreamer(10)
			processedItems := make([]Item, 0)
			stream.Producer(tt.itemsToProduce)
			// Handle Synchronization
			stream.Consumer(&processedItems)
			stream.Close()

			if !reflect.DeepEqual(processedItems, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, processedItems)
			}
		})
	}
}
