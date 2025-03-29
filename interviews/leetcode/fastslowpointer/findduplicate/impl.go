package findduplicate

// findDuplicate is used fast slow pattern to fix the issue.
// In this problem, the array nums can be interpreted as a linked list where the value at each index points to the next index.
func findDuplicate(nums []int) int {
	// slow and fast start from the same position
	slow, fast := nums[0], nums[0]

	// Phase 1: Detect the intersection point of the cycle.
	// Use two pointers, slow and fast, to find a meeting point in the “cycle”.
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	// Phase 2: Find the entrance to the cycle, it means find the duplicate element.
	// Reset the slow pointer, to the start of the array and keep the fast pointer at the meeting point.
	// Move both pointers one step at a time. The point where they meet is the duplicate number.
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
