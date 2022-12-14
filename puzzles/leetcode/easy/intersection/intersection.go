package intersection

func intersection(a, b []int) (result []int) {
	counter := make(map[int]int)
	for _, elem := range a {
		counter[elem]++
	}

	for _, elem := range b {
		if count, ok := counter[elem]; ok && count > 0 {
			counter[elem] -= 1
			result = append(result, elem)
		}
	}

	return
}
