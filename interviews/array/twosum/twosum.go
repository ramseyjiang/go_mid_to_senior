package twosum

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
	record := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ok := record[target-nums[i]]; ok {
			res = append(res, i, val)
			break
		}
		record[nums[i]] = i
	}

	return res
}
