package sample

type IntIterator interface {
	Next() (int, bool)
}

type SliceIntIterator struct {
	slice []int
	index int
}

func (i *SliceIntIterator) Next() (int, bool) {
	if i.index >= len(i.slice) {
		return 0, false
	}
	value := i.slice[i.index]
	i.index++
	return value, true
}

func NewSliceIntIterator(slice []int) IntIterator {
	return &SliceIntIterator{slice: slice}
}
