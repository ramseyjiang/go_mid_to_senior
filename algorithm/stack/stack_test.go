package stack

import (
	"math"
	"reflect"
	"testing"
)

func TestMinStack(t *testing.T) {
	ms := New()

	// Test stack Max int
	wantMax := math.MaxInt64
	// fmt.Println("maxInt64 is ", ms.GetMin())
	t.Run("Test get stack after pop", func(t *testing.T) {
		if got := ms.GetMin(); !reflect.DeepEqual(got, wantMax) {
			t.Errorf("execute() = %v, want %v", got, wantMax)
		}
	})

	// Test stack content
	ms.Push(20)
	ms.Push(39)
	// fmt.Println("MinStack content after push 39 is: ", ms.stack)
	wantStack := []int{20, 39}
	t.Run("Test stack str", func(t *testing.T) {
		if got := ms.stack; !reflect.DeepEqual(got, wantStack) {
			t.Errorf("execute() = %v, want %v", got, wantStack)
		}
	})

	// Test get minimum
	wantMin := 20
	// fmt.Println("MinStack after push 39, the min is: ", ms.GetMin())
	t.Run("Test get minimum", func(t *testing.T) {
		if got := ms.GetMin(); !reflect.DeepEqual(got, wantMin) {
			t.Errorf("execute() = %v, want %v", got, wantMin)
		}
	})

	// Test get stack and minimum
	wantMin1 := 10
	wantStack1 := []int{20, 39, 10}
	ms.Push(10)
	// fmt.Println("MinStack content after push 10 is: ", ms.stack)
	// fmt.Println("MinStack after push 10, the min is: ", ms.GetMin())
	t.Run("Test get stack and minimum", func(t *testing.T) {
		if got := ms.stack; !reflect.DeepEqual(got, wantStack1) {
			t.Errorf("execute() = %v, want %v", got, wantStack1)
		}

		if got := ms.GetMin(); !reflect.DeepEqual(got, wantMin1) {
			t.Errorf("execute() = %v, want %v", got, wantMin1)
		}
	})

	// Test get last element and stack content before pop
	wantLastElement := 10
	wantStackBeforePop := []int{20, 39, 10}
	// fmt.Println(ms.GetLastElement())
	// fmt.Println("MinStack content before pop is: ", ms.stack)
	t.Run("Test get last element and stack content before pop", func(t *testing.T) {
		if got := ms.GetLastElement(); !reflect.DeepEqual(got, wantLastElement) {
			t.Errorf("execute() = %v, want %v", got, wantLastElement)
		}

		if got := ms.stack; !reflect.DeepEqual(got, wantStackBeforePop) {
			t.Errorf("execute() = %v, want %v", got, wantStackBeforePop)
		}
	})

	// Test stack content after pop
	ms.Pop()
	wantStackAfterPop := []int{20, 39}
	t.Run("Test stack content after pop", func(t *testing.T) {
		if got := ms.stack; !reflect.DeepEqual(got, wantStackAfterPop) {
			t.Errorf("execute() = %v, want %v", got, wantStackAfterPop)
		}
	})
	// fmt.Println("MinStack content after pop is: ", ms.stack)

	for i := 0; i <= len(ms.stack); i++ {
		if ms.IsEmpty() == false {
			ms.Pop()
		}
	}
	wantEmptyStack := make([]int, 0)
	t.Run("Test IsEmpty() stack", func(t *testing.T) {
		if got := ms.stack; !reflect.DeepEqual(got, wantEmptyStack) {
			t.Errorf("execute() = %v, want %v", got, wantEmptyStack)
		}
	})
}
