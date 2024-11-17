package twosum

type TwoSum struct {
	hashmap map[int]int
}

func Constructor() TwoSum {
	return TwoSum{make(map[int]int)}
}

func (ts *TwoSum) Add(number int) {
	ts.hashmap[number] += 1
}

func (ts *TwoSum) Find(value int) bool {
	for k, v := range ts.hashmap {
		if _, ok := ts.hashmap[value-k]; ok {
			// if v <= 1 means, the arr only has one value, so even k == value - k, it won't return true, because only one value.
			if k == value-k && v <= 1 {
				continue
			}
			return true
		}
	}
	return false
}
