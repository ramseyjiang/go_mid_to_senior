package rotate

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n // Ensure k is within the bounds of nums length
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[(i+k)%n] = nums[i]
	}

	// Copy the temporary 'result' array back to 'nums'
	copy(nums, result)
}

func rotate2(nums []int, k int) {
	tmp := make([]int, len(nums)-1)
	for i := 1; i <= k; i++ {
		tmpLast := nums[len(nums)-1]
		front := nums[0 : len(nums)-1]
		tmp = append([]int{tmpLast}, front...)
		copy(nums, tmp)
	}
}
