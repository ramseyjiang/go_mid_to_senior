package singlenum

import "sort"

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

func singleNumber2(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}