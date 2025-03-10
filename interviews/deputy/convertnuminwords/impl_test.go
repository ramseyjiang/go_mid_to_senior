package convertnuminwords

import "testing"

func TestConvertNumInWords(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{"three", 3, "Three"},
		{"eight", 8, "Eight"},
		{"nineteen", 19, "Nineteen"},
		{"seventy eight", 78, "Seventy Eight"},
		{"eighty eight", 88, "Eighty Eight"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertNumInWords(tt.num); got != tt.want {
				t.Errorf("ConvertNumInWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
