package twosumlessthank

import "testing"

func TestTwoSumLessThanK(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "Test 1",
			nums: []int{34, 23, 1, 24, 75, 33, 54, 8},
			k:    60,
			want: 58,
		},
		{
			name: "Test 2",
			nums: []int{10, 20, 30},
			k:    25,
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSumLessThanK(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("twoSumLessThanK() = %v, want %v", got, tt.want)
			}
		})
	}
}
