package containsduplicate

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func containsDuplicate1(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]^nums[i+1] == 0 { // Using XOR, the same element is 0, the different element is 1
			return true
		}
	}
	return false
}

func containsDuplicate2(nums []int) bool {
	seen := make(map[int]bool)
	for _, v := range nums {
		if _, exists := seen[v]; exists {
			return true
		}
		seen[v] = false
	}
	return false
}
