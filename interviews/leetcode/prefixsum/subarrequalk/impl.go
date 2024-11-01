package subarrequalk

// SubarraySum returns the number of contiguous subarrays whose sum equals k.
func SubarraySum(nums []int, k int) int {
	prefixSumMap := make(map[int]int)
	prefixSumMap[0] = 1 // Initialize with sum 0 to count subarrays that start from the beginning.
	cumulativeSum, count := 0, 0

	for _, num := range nums {
		cumulativeSum += num

		// Check if there exists a prefix sum that we can subtract to get sum k.
		if occurrences, found := prefixSumMap[cumulativeSum-k]; found {
			count += occurrences
		}

		// Store or update the current cumulative sum in the map.
		prefixSumMap[cumulativeSum]++
	}

	return count
}
