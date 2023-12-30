package servicescb

import (
	"testing"
	"time"

	"gopkg.in/errgo.v2/errors"
)

func TestCircuitBreaker(t *testing.T) {
	failingServiceB := func() error {
		return errors.New("service B failure")
	}

	tests := []struct {
		name            string
		serviceFunc     ServiceFunc
		failCalls       int
		recoveryTimeout time.Duration
		expectError     bool
	}{
		{
			name:            "ServiceA Success",
			serviceFunc:     ServiceA,
			failCalls:       0,
			recoveryTimeout: 1 * time.Second,
			expectError:     false,
		},
		{
			name:            "ServiceB Failure",
			serviceFunc:     failingServiceB, // Use the failing version of ServiceB
			failCalls:       3,
			recoveryTimeout: 1 * time.Second,
			expectError:     true,
		},
		{
			name:            "ServiceA Recovery",
			serviceFunc:     ServiceA,
			failCalls:       2,
			recoveryTimeout: 100 * time.Millisecond,
			expectError:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cb := NewCircuitBreaker(2, tt.recoveryTimeout)

			for i := 0; i < tt.failCalls; i++ {
				cb.Call(tt.serviceFunc)
			}

			time.Sleep(tt.recoveryTimeout)

			err := cb.Call(tt.serviceFunc)
			if (err != nil) != tt.expectError {
				t.Errorf("Test '%s' failed: expected error = %v, got error = %v", tt.name, tt.expectError, err != nil)
			}
		})
	}
}
