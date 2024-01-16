package loggerratelimiter

import "testing"

func TestShouldPrintMessage(t *testing.T) {
	logger := Constructor()

	tests := []struct {
		timestamp int
		message   string
		expected  bool
	}{
		{1, "foo", true},
		{2, "bar", true},
		{3, "foo", false},
		{8, "bar", false},
		{10, "foo", false},
		{11, "foo", true},
		// Additional test cases can be added here.
	}

	for _, tt := range tests {
		t.Run(tt.message, func(t *testing.T) {
			got := logger.ShouldPrintMessage(tt.timestamp, tt.message)
			if got != tt.expected {
				t.Errorf("ShouldPrintMessage(%d, %s) = %v; want %v", tt.timestamp, tt.message, got, tt.expected)
			}
		})
	}
}
