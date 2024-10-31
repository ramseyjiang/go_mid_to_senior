package contiguousarray

// FindMaxLength finds the maximum length of a contiguous subarray with equal numbers of 0 and 1.
func FindMaxLength(nums []int) int {
	// Initialize a map to store the cumulative sum and its first occurrence index.
	sumIndexMap := make(map[int]int)
	// Set initial cumulative sum and max length.
	cumulativeSum, maxLength := 0, 0

	// Set the base case for sumIndexMap to handle the scenario when a subarray from the beginning has a zero cumulative sum.
	sumIndexMap[0] = -1

	// Iterate over the array
	for i, num := range nums {
		// Treat 0 as -1 and 1 as +1 to track the balance.
		if num == 0 {
			cumulativeSum--
		} else {
			cumulativeSum++
		}

		// Check if the cumulative sum has been seen before
		if prevIndex, found := sumIndexMap[cumulativeSum]; found {
			// Update max length if the length of the i - prevIndex is longer
			maxLength = max(maxLength, i-prevIndex)
		} else {
			// Store the cumulative sum with its index if it's seen for the first time
			sumIndexMap[cumulativeSum] = i
		}
	}

	return maxLength
}
