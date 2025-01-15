package nearestsmallernum

// NearestSmallerValues is Time Complexity: O(n), and Space Complexity: O(n)
func NearestSmallerValues(arr []int) []int {
	// Each element is pushed and popped from the stack exactly once, resulting in linear time complexity.
	// In the worst case, the stack could hold all indices from the array, such as in a strictly increasing array.

	// Stack to store indices
	stack := make([]int, 0)
	result := make([]int, len(arr))

	for i, num := range arr {
		// Remove elements from the stack that are greater or equal to the current num,
		// Stack is the LIFO policy, and the stack[:len(stack)-1] is used to remove the last in the indices.
		// The stack provides O(1) operations for push, pop, and peek.
		// This ensures that each element is processed only once, leading to O(n) time complexity.
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= num {
			// If the current element is smaller than elements in the stack, those elements are popped.
			// This is a natural way to find the closest smaller element while eliminating unnecessary larger elements.
			stack = stack[:len(stack)-1]
		}

		// If stack is empty, no smaller value exists
		if len(stack) == 0 {
			result[i] = -1
		} else {
			// Top of the stack is the nearest smaller value
			result[i] = arr[stack[len(stack)-1]]
		}

		stack = append(stack, i)
	}

	return result
}

// The NearestSmallerValuesBruteForce is the brute-force approach,
// It iterates over all previous elements for each element in the array to find the nearest smaller value.
// Time Complexity: O(n²), Each element compares with all previous elements in the worst case.
// Space Complexity: O(1), No extra space is used beyond the result array.
func NearestSmallerValuesBruteForce(arr []int) []int {
	result := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		result[i] = -1
		for j := i - 1; j >= 0; j-- {
			if arr[j] < arr[i] {
				result[i] = arr[j]
				break
			}
		}
	}

	return result
}
