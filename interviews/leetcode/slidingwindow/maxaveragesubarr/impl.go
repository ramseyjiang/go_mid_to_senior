package maxaveragesubarr

func findMaxAverageArr(nums []int, k int) float64 {
	if len(nums) < k {
		return 0 // handle case where nums has fewer than k elements
	}

	// Step 1: Initialize maxSum and sum, calculate the sum of the first k elements
	maxSum, sum := 0, 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum = sum

	// Step 2: Slide the window over the array
	for i := k; i < len(nums); i++ {
		// Subtract the element that is leaving the window and add the element that is entering
		sum = sum - nums[i-k] + nums[i]
		// Track the maximum sum encountered
		maxSum = max(maxSum, sum)
	}

	// Step 3: Calculate and return the maximum average
	return float64(maxSum) / float64(k)
}
