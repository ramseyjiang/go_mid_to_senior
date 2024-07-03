package linkedlistcycle

import "testing"

var tests = []struct {
	name     string
	values   []int
	pos      int
	expected bool
}{
	{
		name:     "Test 1",
		values:   []int{3, 2, 0, -4},
		pos:      1,
		expected: true,
	},
	{
		name:     "Test 2",
		values:   []int{1, 2},
		pos:      0,
		expected: true,
	},
	{
		name:     "Test 3",
		values:   []int{1},
		pos:      -1,
		expected: false,
	},
}

func TestHasCycleIterative(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createCycleList(tt.values, tt.pos)
			result := HasCycleIterative(head)
			if result != tt.expected {
				t.Errorf("HasCycleIterative() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestHasCycleRecursive(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createCycleList(tt.values, tt.pos)
			result := HasCycleRecursive(head)
			if result != tt.expected {
				t.Errorf("HasCycleRecursive() = %v, want %v", result, tt.expected)
			}
		})
	}
}
