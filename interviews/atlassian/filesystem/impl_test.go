package filesystem

import (
	"testing"
)

func TestFileSystem(t *testing.T) {
	tests := []struct {
		name     string
		actions  []string
		paths    []string
		values   []int
		expected []int
	}{
		{
			name:     "Example 1",
			actions:  []string{"CreatePath", "Get"},
			paths:    []string{"/a", "/a"},
			values:   []int{1, 0},
			expected: []int{1, 1},
		},
		{
			name:     "Example 2",
			actions:  []string{"CreatePath", "CreatePath", "Get", "CreatePath", "Get"},
			paths:    []string{"/leet", "/leet/code", "/leet/code", "/c/d", "/c"},
			values:   []int{1, 2, 0, 1, 0},
			expected: []int{1, 1, 2, 0, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := NewFileSystem()
			for i, action := range tt.actions {
				var got int
				if action == "CreatePath" {
					result := fs.CreatePath(tt.paths[i], tt.values[i])
					if result {
						got = 1
					} else {
						got = 0
					}
				}

				if action == "Get" {
					got = fs.Get(tt.paths[i])
				}

				if got != tt.expected[i] {
					t.Errorf("%s(%s, %d) = %d; want %d", action, tt.paths[i], tt.values[i], got, tt.expected[i])
				}
			}
		})
	}
}
