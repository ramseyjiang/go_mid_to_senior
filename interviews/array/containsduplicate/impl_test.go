package containsduplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"Example 1", []int{1, 2, 3, 1}, true},
		{"Example 2", []int{1, 2, 3, 4}, false},
		{"Example 3", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsDuplicate(tt.nums)
			if got != tt.want {
				t.Errorf("containsDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
