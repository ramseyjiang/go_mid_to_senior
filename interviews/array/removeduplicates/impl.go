package removeduplicates

import (
	"golang.org/x/exp/slices"
)

// removeDuplicates is the best way.
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Index for placing the next unique element
	writeIndex := 1

	// Start from the second element in the array and compare with the previous element
	for readIndex := 1; readIndex < len(nums); readIndex++ {
		// If the current element is not the same as the previous one,
		// it is unique and should be moved to the writeIndex.
		if nums[readIndex] != nums[readIndex-1] {
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		}
	}

	// Return the length of the array without duplicates
	return writeIndex
}

// removeDuplicates2 outputs also work, but not pass the leetcode testcases, don't know why.
// Remember the empty struct{}{} uses 0 bytes.
func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tmp := make(map[int]struct{})
	var results []int
	for _, v := range nums {
		tmp[v] = struct{}{} // The empty struct{}{} uses 0 bytes of storage, making it efficient.
		if !slices.Contains(results, v) {
			results = append(results, v)
		}
	}
	nums = results

	return len(tmp)
}
