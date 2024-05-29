package delnode

// ListNode is the definition for a singly linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// DeleteNode deletes the given node from the singly linked list.
func DeleteNode(node *ListNode) {
	if node == nil || node.Next == nil {
		return
	}
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
