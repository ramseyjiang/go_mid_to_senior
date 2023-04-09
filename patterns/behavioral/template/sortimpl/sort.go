package sortimpl

type baseAlgorithm struct {
	intSlice []int
}

func (m baseAlgorithm) Len() int {
	return len(m.intSlice)
}

func (m baseAlgorithm) Swap(i, j int) {
	m.intSlice[i], m.intSlice[j] = m.intSlice[j], m.intSlice[i]
}

func (m baseAlgorithm) Less(i, j int) bool {
	return m.intSlice[i] < m.intSlice[j]
}
