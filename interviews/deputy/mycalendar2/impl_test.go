package mycalendar2

import "testing"

func TestMyCalendarTwo_Book(t *testing.T) {
	tests := []struct {
		name   string
		inputs [][2]int
		wants  []bool
	}{
		{
			name:   "basic triple book check",
			inputs: [][2]int{{10, 20}, {15, 25}, {10, 20}},
			wants:  []bool{true, true, false},
		},
		{
			name:   "multiple area not overlaps check",
			inputs: [][2]int{{10, 20}, {50, 60}, {10, 40}, {5, 15}},
			wants:  []bool{true, true, true, false},
		},
		// 其他测试用例...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cal := Constructor()
			for i := range tt.inputs {
				got := cal.Book(tt.inputs[i][0], tt.inputs[i][1])
				if got != tt.wants[i] {
					t.Errorf("Book(%v) = %v, want %v", tt.inputs[i], got, tt.wants[i])
				}
			}
		})
	}
}
