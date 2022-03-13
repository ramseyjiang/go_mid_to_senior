package viperpkg

import "testing"

// go test -v -run TestEntry
func TestEntry(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "env.toml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Entry()
		})
	}
}
