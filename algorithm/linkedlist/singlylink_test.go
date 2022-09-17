package linkedlist

import (
	"reflect"
	"testing"
)

func TestSinglyLink_InsertAtHead(t *testing.T) {
	link := NewSinglyLinkedList[int]()
	link.InsertAtHead(1)
	link.InsertAtHead(2)
	link.InsertAtHead(3)
	link.InsertAtHead(4)

	expected := []int{4, 3, 2, 1}
	values := link.Values()
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestSinglyLink_InsertAtTail(t *testing.T) {
	link := NewSinglyLinkedList[int]()
	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)
	link.InsertAtTail(4)

	expected := []int{1, 2, 3, 4}
	values := link.Values()
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestSinglyLink_InsertAt(t *testing.T) {
	link := NewSinglyLinkedList[int]()
	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)
	link.InsertAtTail(4)

	_ = link.InsertAt(0, 0)
	_ = link.InsertAt(1, 11)
	_ = link.InsertAt(3, 6)

	expected := []int{0, 11, 1, 6, 2, 3, 4}
	values := link.Values()
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestSinglyLink_DeleteAt(t *testing.T) {
	link := NewSinglyLinkedList[int]()
	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)
	link.InsertAtTail(4)

	_ = link.DeleteAtHead()
	_ = link.DeleteAt(2)

	expected := []int{2, 3}
	values := link.Values()
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}

}
