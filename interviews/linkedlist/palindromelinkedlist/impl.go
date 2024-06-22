package palindromelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindromeIterative(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// Find the middle of the linked list
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse the second half of the linked list
	var prev *ListNode
	curr := slow
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}

	// Check if the first half and the reversed second half are the same
	firstHalf, secondHalf := head, prev
	for secondHalf != nil {
		if firstHalf.Val != secondHalf.Val {
			return false
		}
		firstHalf = firstHalf.Next
		secondHalf = secondHalf.Next
	}

	return true
}

func isPalindromeRecursive(head *ListNode) bool {
	frontPointer := head

	var recursivelyCheck func(*ListNode) bool
	recursivelyCheck = func(current *ListNode) bool {
		if current != nil {
			if !recursivelyCheck(current.Next) {
				return false
			}
			if current.Val != frontPointer.Val {
				return false
			}
			frontPointer = frontPointer.Next
		}
		return true
	}

	return recursivelyCheck(head)
}
