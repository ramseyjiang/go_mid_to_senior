package majorityelement

import "fmt"

func FindMajority(num int, arr []int) int {
	tmp := make(map[int]int)
	// tmp[v]++ simplifies the incrementation logic by removing the need to check if the key exists,
	// since Go automatically initializes missing int map values to 0.
	for _, v := range arr {
		// if _, exists := tmp[v]; exists {
		tmp[v]++
		// } else {
		// 	tmp[v] = 1
		// }
	}

	for _, value := range tmp {
		if value > num/2 {
			return value
		}
	}

	return -1
}

func FindMajority2(n int, nums []int) int {
	count := 0
	candidate := 0

	// Boyer-Moore Voting Algorithm
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	fmt.Println(count)
	// Verify if the candidate is actually a majority element
	count = 0
	for _, num := range nums {
		if num == candidate {
			count++
		}
	}
	fmt.Println(count)
	if count > len(nums)/2 {
		return candidate
	}
	return -1
}
