package movezeros

// first pointer for control left zero
// second pointer for iterating
func moveZeroes(nums []int) {
	zeroIndex := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[zeroIndex] = nums[zeroIndex], nums[i]
			zeroIndex++
		}
	}
}

func moveZeroes2(nums []int) {
	insertPos := 0

	// move all numbers gather than 0 to front
	for _, num := range nums {
		if num != 0 {
			nums[insertPos] = num
			insertPos++
		}
	}

	// fill all last positions 0
	for insertPos < len(nums) {
		nums[insertPos] = 0
		insertPos++
	}
}
