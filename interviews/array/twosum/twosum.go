package twosum

import (
	"slices"
	"sort"
)

// two Sum2 is used the bubble sort way to solve it.
func twoSum2(nums []int, target int) (res []int) {
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] == target-nums[i] {
				res = []int{i, j}
			}
		}
	}
	return
}

func twoSum(nums []int, target int) (res []int) {
	record := make(map[int]int, 2)
	for i := 0; i < len(nums); i++ {
		if val, ok := record[target-nums[i]]; ok {
			res = append(res, i, val)
			break
		}
		record[nums[i]] = i
	}

	return res
}

// Using two pointer pattern to solve, this way should update unit test then it can pass tests.
func twoSum3(nums []int, target int) []int {
	length := len(nums)
	sum, l, r := 0, 0, length-1

	// copy the nums to tmpNums, use tmpNums to confirm the nums index,
	// as the slice is not sorted before, after sorted, the original index is changed, that's why copy it.
	tmpNums := make([]int, length)
	copy(tmpNums, nums)
	tmpRes, res := make([]int, length), make([]int, length)
	sort.Ints(nums)

	for l < r {
		sum = nums[l] + nums[r]
		if sum == target {
			tmpRes = append(tmpRes, l, r)
			break
		} else if sum < target {
			l++
		} else {
			r--
		}
	}

	res = []int{slices.Index(tmpNums, nums[l]), slices.Index(tmpNums, nums[r])}

	return res
}
