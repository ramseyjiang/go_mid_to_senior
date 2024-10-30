package rangesumquery

// NumArray struct will hold the nums array and the prefix sum array
type NumArray struct {
	prefixSum []int
}

// PrefixSum to initialize the PrefixSum object with the nums array
//
//	(*NumArray).
//
// This allows you to modify the NumArray directly via the pointer.
// Using pointers avoids the need to duplicate the entire object,
// which is more space-efficient, especially when dealing with larger data structures.
func PrefixSum(nums []int) *NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	return &NumArray{nums}
}

// The first function that returns a pointer (*NumArray) uses slightly less space
// because it returns a reference to the NumArray struct rather than making a full copy of the struct.
// In contrast, the second function that returns a value (NumArray) creates a full copy of the struct, which takes more memory.

// Returns a value type (NumArray), meaning a copy of the struct is returned.
// Any changes made to the struct outside the function won’t reflect back unless explicitly returned or updated.
// func PrefixSum(nums []int) NumArray {
// 	for i := 1; i < len(nums); i++ {
// 		nums[i] += nums[i-1]
// 	}
//
// 	return NumArray{nums}
// }

// SumRange function returns the sum of elements between left and right inclusive
func (n *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return n.prefixSum[right]
	}

	return n.prefixSum[right] - n.prefixSum[left-1]
}
