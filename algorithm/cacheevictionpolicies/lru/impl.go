package lru

import (
	"container/list"
)

// CacheLRU represents a Least Recent Used cache
type CacheLRU struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

// Pair stores key-value pairs
type Pair struct {
	key   int
	value int
}

// NewCacheLRU initializes the CacheLRU with a given capacity
func NewCacheLRU(capacity int) *CacheLRU {
	return &CacheLRU{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

// Get retrieves a value from the cache
func (c *CacheLRU) Get(key int) int {
	if elem, found := c.cache[key]; found {
		// Move the accessed element to the front (most recently used)
		c.list.MoveToFront(elem)
		return elem.Value.(*Pair).value
	}
	return -1 // Key not found
}

// Put adds a key-value pair to the cache
func (c *CacheLRU) Put(key int, value int) {
	if elem, found := c.cache[key]; found {
		// Update the value and move the element to the front
		elem.Value.(*Pair).value = value
		c.list.MoveToFront(elem)
		return
	}

	// Add new key-value pair
	if c.list.Len() >= c.capacity {
		// Evict the least recently used element
		lru := c.list.Back()
		if lru != nil {
			delete(c.cache, lru.Value.(*Pair).key)
			c.list.Remove(lru)
		}
	}

	pair := &Pair{key: key, value: value}
	elem := c.list.PushFront(pair)
	c.cache[key] = elem
}
