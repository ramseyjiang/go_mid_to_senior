package stack

import "math"

// MinStack is a simple struct with two fields. It is a stack with a slice of integers and contains its minimum value.
// A stack always follows an order Last-In-First-Out (LIFO).
// The operations to add or take from the stack are called push and pop respectively.
type MinStack struct {
	stack []int
	min   int
}

func New() MinStack {
	maxInt64 := math.MaxInt64
	return MinStack{min: maxInt64}
}

// Push method is used to push an element, if it is lower than the current min
func (ms *MinStack) Push(item int) {
	ms.stack = append(ms.stack, item)
	if item < ms.min {
		ms.min = item
	}
}

// Pop is used to remove the last element from the stack and save its value in variable last
func (ms *MinStack) Pop() {
	if ms.IsEmpty() {
		return
	}
	length := len(ms.stack)
	last := ms.stack[length-1]
	ms.stack = ms.stack[:length-1]

	if last != ms.min {
		return
	}

	// popped an element off that had the min value, find the minimum by iterating through the list.
	ms.min = math.MaxInt64
	for _, x := range ms.stack {
		if x < ms.min {
			ms.min = x
		}
	}
}

func (ms *MinStack) IsEmpty() bool {
	if len(ms.stack) == 0 {
		return true
	}
	return false
}

func (ms *MinStack) GetLastElement() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.min
}
