package plusone

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{"Example 1", []int{1, 2, 3}, []int{1, 2, 4}},
		{"Example 2", []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{"Example 3", []int{9}, []int{1, 0}},
		{"Example 4", []int{9, 9}, []int{1, 0, 0}},
		{"Example 5", []int{8, 9}, []int{9, 0}},
		// Additional test cases can be added here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plusOne(tt.digits)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne(%v) = %v, want %v", tt.digits, got, tt.want)
			}
		})
	}
}
