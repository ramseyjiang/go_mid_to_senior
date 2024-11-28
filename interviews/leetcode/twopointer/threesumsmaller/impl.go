package threeSumSmaller

import "sort"

// threeSumSmaller returns the count of triplets where the sum is less than the target.
func threeSumSmaller(nums []int, target int) int {
	count := 0
	n := len(nums)

	// Sort the array
	sort.Ints(nums)

	// Iterate over the array
	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1

		// Two-pointer approach
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < target {
				// If the sum is less than the target, all triplets from left to right are valid
				count += right - left
				left++
			} else {
				right--
			}
		}
	}

	return count
}
