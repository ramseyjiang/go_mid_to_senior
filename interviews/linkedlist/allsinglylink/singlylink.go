package allsinglylink

import "fmt"

// Singly Linked List is a struct with a head pointer to the first LinkNode, and a length property to record the size of the list.

type LinkNode[T any] struct {
	Value T
	Next  *LinkNode[T]
}

func NewLinkNode[T any](value T) *LinkNode[T] {
	return &LinkNode[T]{
		Value: value,
		Next:  nil,
	}
}

// SinglyLinkedList is a singly linked list
type SinglyLinkedList[T any] struct {
	Head   *LinkNode[T]
	length int
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		Head:   nil,
		length: 0,
	}
}

// Get the value of the index node in the linked list, if the index is valid.
func (l *SinglyLinkedList[T]) Get(index int) (T, bool) {
	// special case: input index is out of range, return error
	var t T
	if index < 0 || index > l.length-1 {
		return t, false
	}
	// general case: find the node at the index, then return its value
	cur := l.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Value, true
}

// InsertAtHead Insert a new node to the head of list
func (l *SinglyLinkedList[T]) InsertAtHead(value T) {
	newNode := NewLinkNode(value)
	newNode.Next = l.Head
	l.Head = newNode
	l.length++
}

// InsertAtTail Append a new node to the tail of list, traverse the list until we find the last node and point it to the new node.
func (l *SinglyLinkedList[T]) InsertAtTail(value T) {
	// special case: the list is empty
	if l.Head == nil {
		l.InsertAtHead(value)
		return
	}

	// general case: find the last node of list, then append a new node to the tail
	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	newNode := NewLinkNode(value)
	cur.Next = newNode
	l.length++ // update the length of list
}

// InsertAt Add a node of value val before the index node in the linked list.
// If index equals the length of the linked list, the node will be appended to the end of the linked list.
// If index is greater than the length, the node will not be inserted.
func (l *SinglyLinkedList[T]) InsertAt(index int, value T) error {
	// special case: input index is out of range, return error
	size := l.length
	if index < 0 || index > size {
		return fmt.Errorf("index out of range")
	}

	// special case: index is 0, insert at head by use InsertAtHead method
	if index == 0 {
		l.InsertAtHead(value)
		return nil
	}

	// special case: index equals the length of list, insert at tail by use InsertAtTail method
	if index == size {
		l.InsertAtTail(value)
		return nil
	}

	// general case: find the node before the index, then insert a new node after the found node
	cur := l.Head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}
	newNode := NewLinkNode(value)
	newNode.Next = cur.Next
	cur.Next = newNode
	l.length++

	return nil
}

// DeleteAtHead Delete the head node of the list is straightforward
func (l *SinglyLinkedList[T]) DeleteAtHead() error {
	if l.Head == nil {
		return fmt.Errorf("list is empty")
	}

	l.Head = l.Head.Next
	l.length--
	return nil
}

// DeleteAtTail Delete the tail node of the list.
func (l *SinglyLinkedList[T]) DeleteAtTail() error {
	if l.Head == nil {
		return fmt.Errorf("list is empty")
	}
	// special case: the list has only one node
	if l.Head.Next == nil {
		return l.DeleteAtHead()
	}

	// general case: find the second last node of list, then delete the last node, after that the second last node is the last node.
	cur := l.Head
	for cur.Next.Next != nil {
		cur = cur.Next
	}
	cur.Next = nil
	l.length--
	return nil
}

// DeleteAt Delete the index node in the linked list, if the index is valid.
func (l *SinglyLinkedList[T]) DeleteAt(index int) error {
	size := l.length
	// special case: input index is out of range, return error
	if index < 0 || index > size-1 {
		return fmt.Errorf("index out of range")
	}
	// special case: index is 0, delete at head by use DeleteAtHead method
	if index == 0 {
		return l.DeleteAtHead()
	}

	// general case: find the node before the index, then delete its next node
	cur := l.Head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	l.length--
	return nil

}

func (l *SinglyLinkedList[T]) Length() int {
	return l.length
}

func (l *SinglyLinkedList[T]) Values() []T {
	var values []T
	cur := l.Head
	for cur != nil {
		values = append(values, cur.Value)
		cur = cur.Next
	}
	return values
}
