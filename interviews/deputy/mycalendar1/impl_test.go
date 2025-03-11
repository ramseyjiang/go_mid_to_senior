package mycalendar1

import "testing"

func TestMyCalendarBook(t *testing.T) {
	tests := []struct {
		name     string
		bookings [][]int // [start, end] pairs
		want     []bool  // expected results
	}{
		{
			name:     "non-overlapping bookings",
			bookings: [][]int{{10, 20}, {25, 30}, {5, 8}},
			want:     []bool{true, true, true},
		},
		{
			name:     "overlapping bookings",
			bookings: [][]int{{10, 20}, {15, 25}, {20, 30}},
			want:     []bool{true, false, true},
		},
		{
			name:     "exact overlap",
			bookings: [][]int{{5, 10}, {5, 10}},
			want:     []bool{true, false},
		},
		{
			name:     "edge cases",
			bookings: [][]int{{5, 10}, {10, 20}, {9, 11}},
			want:     []bool{true, true, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Constructor()
			for i, booking := range tt.bookings {
				got := c.Book(booking[0], booking[1])
				if got != tt.want[i] {
					t.Errorf("Book(%d, %d) = %v, want %v",
						booking[0], booking[1], got, tt.want[i])
				}
			}
		})
	}
}
