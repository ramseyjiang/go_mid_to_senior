package twosumlessthank

import "sort"

func twoSumLessThanK(nums []int, k int) int {
	length := len(nums)
	sort.Ints(nums) // sorted the nums by ascending order
	left, right, sum, maxSum := 0, length-1, 0, 0

	for left < right {
		sum = nums[left] + nums[right]
		// If sum is less than k, move the left pointer to right.
		// If sum is gather than k, move the right pointer to left.
		if sum < k {
			maxSum = max(sum, maxSum)
			left++
		} else {
			right--
		}
	}

	if maxSum == 0 {
		return -1
	} else {
		return maxSum
	}
}
