package sample

type IntIterator interface {
	HasNext() bool
	Next() (int, bool)
}

type SliceIntIterator struct {
	slice []int
	index int
}

func (i *SliceIntIterator) HasNext() bool {
	if i.index >= len(i.slice) {
		return true
	}
	return false
}

func (i *SliceIntIterator) Next() (int, bool) {
	if i.HasNext() {
		return 0, false
	}
	value := i.slice[i.index]
	i.index++
	return value, true
}

func NewSliceIntIterator(slice []int) IntIterator {
	return &SliceIntIterator{slice: slice}
}
