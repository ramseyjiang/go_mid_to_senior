package heap

import "sync"

type Heap struct {
	elements []int
	mutex    sync.RWMutex
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

// Push adds a new element to the top of the heap, while pop removes the top of the heap.
// However, unlike a stack, the element that is returned by pop will always be the smallest or the largest element in the heap,
// depending on whether it’s a min heap or a max heap— that’s how the heap works.
func (h *Heap) Push(element int) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	// append to the end of the elements slice
	h.elements = append(h.elements, element)

	// starting from the index of the last element
	// swap the index and the index its parent if it's larger
	i := len(h.elements) - 1
	for ; h.elements[i] > h.elements[parent(i)]; i = parent(i) {
		h.swap(i, parent(i))
	}
}

// Pop means we take top element of the heap out, and we need to reorganise the heap afterwards.
func (h *Heap) Pop() (ele int) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	// take the top element out of the heap
	ele = h.elements[0]

	// take the last element and move it the top of the heap
	h.elements[0] = h.elements[len(h.elements)-1]
	// reorganise h.elements order after pop.
	h.elements = h.elements[:len(h.elements)-1]

	// call the recursive function rearrange the heap, passing it the index of the top of the heap
	h.rearrange(0)
	return
}

func (h *Heap) rearrange(i int) {
	// assume the given index is the largest
	largest := i

	// get the left and right child indices, and the size of the heap
	left, right, size := leftChild(i), rightChild(i), len(h.elements)

	// if left child is larger than the current largest,
	// set the left to be the largest
	if left < size && h.elements[left] > h.elements[largest] {
		largest = left
	}

	// if right child is larger than the current largest,
	// set the right to be the largest
	if right < size && h.elements[right] > h.elements[largest] {
		largest = right
	}

	// if the largest has changed, swap the positions and recurse
	if largest != i {
		h.swap(i, largest)
		h.rearrange(largest)
	}
}

func (h *Heap) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}
