package islandsnum

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	testCases := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "test 1",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			name: "test 2",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
		{
			name: "test empty grid",
			grid: [][]byte{},
			want: 0,
		},
		{
			name: "test no islands",
			grid: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			want: 0,
		},
		{
			name: "test single land cell",
			grid: [][]byte{
				{'1'},
			},
			want: 1,
		},
		{
			name: "test single water cell",
			grid: [][]byte{
				{'0'},
			},
			want: 0,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := numIslands(tt.grid)
			if got != tt.want {
				t.Errorf("numIslands() got = %v, want %v", got, tt.want)
			}
		})
	}
}
