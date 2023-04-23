package sample

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIntegerIterator(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	collection := NewIntegerCollection(items)
	iterator := collection.Iterator()

	expectedItems := []int{1, 2, 3, 4, 5}
	i := 0

	for iterator.HasNext() {
		item := iterator.Next()
		assert.Equal(t, item, expectedItems[i])
		if item != expectedItems[i] {
			t.Errorf("Expected %d, but got %d", expectedItems[i], item)
		}
		i++
	}

	if i != len(expectedItems) {
		t.Errorf("Expected to iterate %d times, but iterated %d times", len(expectedItems), i)
	}
}

func TestEmptyIntegerCollection(t *testing.T) {
	var items []int
	collection := NewIntegerCollection(items)
	iterator := collection.Iterator()

	assert.Equal(t, false, iterator.HasNext())
	if iterator.HasNext() {
		t.Errorf("Expected HasNext() to return false for empty collection")
	}

	assert.Equal(t, nil, iterator.Next())
	if iterator.Next() != nil {
		t.Errorf("Expected Next() to return nil for empty collection")
	}
}
