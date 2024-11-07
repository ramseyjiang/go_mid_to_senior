package minsizesubarrsum

// minSubArrLen finds the minimal length of a contiguous subarray of which the sum ≥ target.
// If there is no such subarray, it returns 0.
func minSubArrLen(target int, nums []int) int {
	n := len(nums)
	minLength := n + 1 // Start with a length greater than any possible subarray
	left, right, currentSum := 0, 0, 0

	for right < n {
		currentSum += nums[right]

		// When sum >= target, try to shrink the window from the left
		for currentSum >= target {
			minLength = min(minLength, right-left+1)
			currentSum -= nums[left]
			left++
		}
		right++
	}

	if minLength == n+1 {
		return 0 // No valid subarray found
	}
	return minLength
}
