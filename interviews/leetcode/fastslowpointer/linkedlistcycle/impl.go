package linkedlistcycle

// ListNode is a node in a singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// HasCycleIterative detects if a cycle exists in the linked list using an iterative approach
func HasCycleIterative(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}

// HasCycleRecursive detects if a cycle exists in the linked list using a single recursive function
func HasCycleRecursive(head *ListNode) bool {
	var detect func(slow, fast *ListNode) bool
	detect = func(slow, fast *ListNode) bool {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
		return detect(slow, fast)
	}
	return detect(head, head)
}
