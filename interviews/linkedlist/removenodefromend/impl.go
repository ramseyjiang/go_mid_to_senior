package removenodefromend

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	first, second := dummy, dummy

	// Move first n+1 steps ahead
	for i := 0; i <= n; i++ {
		first = first.Next
	}

	// Move first to the end, maintaining the gap.
	// Assume the ListNode length is l, after the loop below, the second pointer position is l-n.
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// Skip the desired node
	second.Next = second.Next.Next

	return dummy.Next
}
