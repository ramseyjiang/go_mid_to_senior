package twosum

/*
Runtime: 42 ms, faster than 25.13% of Go online submissions for Two Sum.
Memory Usage: 3.5 MB, less than 95.16% of Go online submissions for Two Sum.
*/
func twoSum1(nums []int, target int) (res []int) {
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] == target-nums[i] {
				res = []int{i, j}
			}
		}
	}
	return
}

/*
Runtime: 8 ms, faster than 82.76% of Go online submissions for Two Sum.
Memory Usage: 4.3 MB, less than 39.15% of Go online submissions for Two Sum.

Runtime: 4 ms, faster than 95.95% of Go online submissions for Two Sum.
Memory Usage: 4.2 MB, less than 56.57% of Go online submissions for Two Sum.
*/
func twoSum2(nums []int, target int) (res []int) {
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
