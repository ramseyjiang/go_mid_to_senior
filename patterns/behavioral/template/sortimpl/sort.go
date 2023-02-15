package sortimpl

// type Interface interface {
// Len() int
// Less(i, j int) bool
// Swap(i, j int)
// }
// The above is the go source code in the sort package. It is using the template pattern.
// In other words, write a type that implements this Interface so that the Sort package can be used to sort any int slice.

type intList []int

// The following methods are used to implement the template interface methods.

func (m intList) Len() int {
	return len(m)
}

func (m intList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m intList) Less(i, j int) bool {
	return m[i] < m[j]
}
