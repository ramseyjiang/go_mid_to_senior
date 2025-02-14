package courseschedule

import "testing"

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{
			name:          "Example 1: Single prerequisite, no cycle",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          true,
		},
		{
			name:          "Example 2: Simple cycle",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			want:          false,
		},
		{
			name:          "Multiple courses, no cycle",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}},
			want:          true,
		},
		{
			name:          "Multiple courses, has cycle",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}, {1, 3}},
			want:          false,
		},
		{
			name:          "No prerequisites at all",
			numCourses:    3,
			prerequisites: [][]int{},
			want:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canFinish(tt.numCourses, tt.prerequisites)
			if got != tt.want {
				t.Errorf("canFinish(%d, %v) = %v; want %v",
					tt.numCourses, tt.prerequisites, got, tt.want)
			}
		})
	}
}
