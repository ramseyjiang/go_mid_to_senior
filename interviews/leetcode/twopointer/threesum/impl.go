package threesum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	n := len(nums)

	// Sort the array
	sort.Ints(nums)

	// Iterate over the array
	for i := 0; i < n-2; i++ {
		// Skip duplicate values for the first element of the triplet
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Two-pointer approach
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				// Found a triplet
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Move both pointers to the next unique elements
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
