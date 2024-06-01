package reverselist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var rev *ListNode
	curr := head

	for curr != nil {
		nextTemp := curr.Next
		curr.Next = rev
		rev = curr
		curr = nextTemp
	}
	return rev
}
