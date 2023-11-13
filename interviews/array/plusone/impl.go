package plusone

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		// Handle all the last digit less than 9 and return digits.
		if digits[i] < 9 {
			digits[i]++
			return digits
		}

		digits[i] = 0
	}

	// Handle the case where an additional digit is needed
	return append([]int{1}, digits...)
}
