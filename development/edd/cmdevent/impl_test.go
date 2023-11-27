package cmdevent

import (
	"strings"
	"testing"
)

// TestCommandEventManager tests the Run method of CommandEventManager.
func TestCommandEventManager(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"No Command", "", false},
		{"Single Command", "test-command", false},
		{"Command With Args", "test-command:arg1 arg2", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager := NewCommandEventManager()

			// Mock listener to capture command invocations
			var invokedCommands []string
			manager.Add("any-command", func(cmd *Command) {
				invokedCommands = append(invokedCommands, cmd.Kind)
			})

			inputReader := strings.NewReader(tt.input)
			manager.Run(inputReader)

			// Test assertions here...
			// For example, check if the invokedCommands slice contains the expected commands
		})
	}
}
