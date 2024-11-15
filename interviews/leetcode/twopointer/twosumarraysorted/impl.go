package twosumarraysorted

func twoSum(nums []int, target int) []int {
	l := len(nums)
	sum, left, right := 0, 0, l-1
	res := make([]int, 0)

	for left < right {
		sum = nums[left] + nums[right]
		if sum == target {
			res = append(res, left+1, right+1)
			break
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return res
}
