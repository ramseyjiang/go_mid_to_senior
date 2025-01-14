package linkedlistcycle

func createCycleList(values []int, pos int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head
	var cycleNode *ListNode

	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
		if i == pos {
			cycleNode = current
		}
	}

	if pos == 0 {
		cycleNode = head
	}

	if pos >= 0 {
		current.Next = cycleNode
	}

	return head
}
