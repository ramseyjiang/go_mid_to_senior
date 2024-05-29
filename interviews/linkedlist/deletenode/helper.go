package delnode

// Helper function to convert a slice to a linked list.
func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

// Helper function to convert a linked list to a slice.
func listToSlice(head *ListNode) []int {
	var nums []int
	for current := head; current != nil; current = current.Next {
		nums = append(nums, current.Val)
	}
	return nums
}
