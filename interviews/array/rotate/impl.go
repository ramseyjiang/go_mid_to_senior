package rotate

func Rotate(nums []int, k int) {
	n := len(nums)
	k %= n // Ensure k is within the bounds of nums length
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[(i+k)%n] = nums[i]
	}

	// Copy the temporary 'result' array back to 'nums'
	copy(nums, result)
}

func Rotate2(nums []int, k int) {
	n := len(nums)
	k %= n // Ensure k is within the bounds of nums length
	tmp := make([]int, n-1)
	for i := 1; i <= k; i++ {
		tmpLast := nums[n-1]
		front := nums[0 : n-1]
		tmp = append([]int{tmpLast}, front...)
		copy(nums, tmp) // The copy is in the loop, it will waste space.
	}
}
