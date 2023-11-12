package intersectiontwoarr

import (
	"reflect"
	"sort"
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{"Example 1", []int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{"Example 2", []int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
		// Additional test cases can be added here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intersect(tt.nums1, tt.nums2)
			sort.Ints(got) // Sorting because the order doesn't matter
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}
