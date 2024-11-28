package threesumclosest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	// Sort the array
	sort.Ints(nums)

	n := len(nums)
	closestSum := math.MaxInt32

	// Iterate over the array
	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1

		for left < right {
			// Calculate the current sum of the triplet
			sum := nums[i] + nums[left] + nums[right]

			// Update the closest sum if needed
			if abs(target-sum) < abs(target-closestSum) {
				closestSum = sum
			}

			// Move pointers based on the comparison of sum and target
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				// If the exact target is hit, return the sum immediately
				return sum
			}
		}
	}

	return closestSum
}

// Helper function to calculate the absolute value of a number
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
