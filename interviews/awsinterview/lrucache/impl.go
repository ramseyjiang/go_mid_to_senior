package lrucache

import "container/list"

// LRUCache represents an LRU cache
type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

// Pair stores key-value pairs
type Pair struct {
	key   int
	value int
}

// Constructor initializes the CacheLRU with a given capacity
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

// Get retrieves a value from the cache
func (c *LRUCache) Get(key int) int {
	if elem, found := c.cache[key]; found {
		// Move the accessed element to the front (most recently used)
		c.list.MoveToFront(elem)
		return elem.Value.(*Pair).value
	}
	return -1
}

// Put adds a key-value pair to the cache
func (c *LRUCache) Put(key int, value int) {
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
