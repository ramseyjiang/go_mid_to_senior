package transaction

import "testing"

func TestProcessLogs(t *testing.T) {
	tests := []struct {
		name      string
		logs      []string
		threshold int32
		want      []string
	}{
		{
			name:      "ExampleTest",
			logs:      []string{"88 99 200", "88 99 300", "99 32 100", "12 12 15"},
			threshold: 2,
			want:      []string{"88", "99"},
		},
		{
			name:      "SingleTransaction",
			logs:      []string{"1 2 100"},
			threshold: 1,
			want:      []string{"1", "2"},
		},
		{
			name:      "SelfTransaction",
			logs:      []string{"1 1 100"},
			threshold: 1,
			want:      []string{"1"},
		},
		{
			name:      "MultipleSelfTransactions",
			logs:      []string{"1 1 100", "1 1 200"},
			threshold: 2,
			want:      []string{"1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := processLogsOptimized(tt.logs, tt.threshold)
			if !equalSlices(got, tt.want) {
				t.Errorf("processLogs() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to check if two string slices are equal.
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
