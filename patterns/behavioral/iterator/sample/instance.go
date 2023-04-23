package sample

// Iterator is defined as the iterator interface. Step 1.
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// Aggregate is defined as the Aggregate interface. Step 2.
type Aggregate interface {
	Iterator() Iterator
}

// IntegerCollection is used to implement the concrete aggregate. Step 4.
type IntegerCollection struct {
	items []int
}

func (c *IntegerCollection) Iterator() Iterator {
	return &IntegerIterator{collection: c, index: -1}
}

func NewIntegerCollection(items []int) *IntegerCollection {
	return &IntegerCollection{items: items}
}

// IntegerIterator is used to implement the concrete iterator.
type IntegerIterator struct {
	collection *IntegerCollection
	index      int
}

func (it *IntegerIterator) HasNext() bool {
	return it.index < len(it.collection.items)-1
}

func (it *IntegerIterator) Next() interface{} {
	if it.HasNext() {
		it.index++
		return it.collection.items[it.index]
	}
	return nil
}
