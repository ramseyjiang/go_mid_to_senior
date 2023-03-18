package sample

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestFlyweight(t *testing.T) {
	factory := &FlyweightFactory{make(map[interface{}]Flyweight)}
	client := &Client{factory}

	// Get flyweight for key1
	flyweight1 := client.factory.GetFlyweight("key1")

	// Get flyweight for key1 again - should return same object
	flyweight2 := client.factory.GetFlyweight("key1")

	assert.Equal(t, flyweight1, flyweight2)
	if flyweight1 != flyweight2 {
		t.Errorf("Expected flyweight1 to equal flyweight2")
	}

	// Get flyweight for key2
	flyweight3 := client.factory.GetFlyweight("key2")

	assert.NotEqual(t, flyweight1, flyweight3)
	if flyweight1 == flyweight3 {
		t.Errorf("Expected flyweight1 to not equal flyweight3")
	}
}
