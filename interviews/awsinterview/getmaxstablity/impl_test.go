package getmaxstablity

import (
	"testing"
)

func TestGetMaxStability(t *testing.T) {
	tests := []struct {
		name         string
		reliability  []int32
		availability []int32
		want         int32
	}{
		{
			name:         "single server",
			reliability:  []int32{5},
			availability: []int32{3},
			want:         15,
		},
		{
			name:         "two servers optimal subset",
			reliability:  []int32{3, 5},
			availability: []int32{2, 1},
			want:         8,
		},
		{
			name:         "three servers max at end",
			reliability:  []int32{2, 3, 1},
			availability: []int32{5, 4, 6},
			want:         24,
		},
		{
			name:         "max in middle of sorted list",
			reliability:  []int32{1, 3, 5},
			availability: []int32{10, 8, 5},
			want:         45,
		},
		{
			name:         "modulo overflow case",
			reliability:  []int32{1000000000},
			availability: []int32{2},
			want:         999999993,
		},
		{
			name:         "empty input",
			reliability:  []int32{},
			availability: []int32{},
			want:         0,
		},
		{
			name:         "same availability all servers",
			reliability:  []int32{5, 5, 5},
			availability: []int32{3, 3, 3},
			want:         45,
		},
		{
			name:         "all zero availability",
			reliability:  []int32{2, 4},
			availability: []int32{0, 0},
			want:         0,
		},
		{
			name:         "mixed zero availability",
			reliability:  []int32{4, 2},
			availability: []int32{5, 0},
			want:         20,
		},
		{
			name:         "large numbers no overflow",
			reliability:  []int32{2147483647, 2147483647},
			availability: []int32{2147483647, 2147483647},
			want:         (2147483647 * (2147483647 * 2)) % 1000000007,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getMaxStability(tt.reliability, tt.availability)
			if got != tt.want {
				t.Errorf("getMaxStability() = %v, want %v", got, tt.want)
			}
		})
	}
}
