package fsm

import (
	"testing"
	"time"
)

func TestCircuitBreaker(t *testing.T) {
	resetTimeout := 100 * time.Millisecond
	cb := NewCircuitBreaker(2, resetTimeout)

	tests := []struct {
		name string
		run  func(t *testing.T)
	}{
		{
			name: "Closed to Open Transition",
			run: func(t *testing.T) {
				cb.Failure()
				cb.Failure()
				if cb.State() != Open {
					t.Errorf("expected state to be Open, got %s", cb.State())
				}
			},
		},
		{
			name: "Open to HalfOpen Transition",
			run: func(t *testing.T) {
				time.Sleep(resetTimeout + 10*time.Millisecond) // Wait for resetTimeout to expire
				if cb.State() != HalfOpen {
					t.Errorf("expected state to be HalfOpen, got %s", cb.State())
				}
			},
		},
		{
			name: "HalfOpen to Closed Transition",
			run: func(t *testing.T) {
				cb.Success()
				if cb.State() != Closed {
					t.Errorf("expected state to be Closed, got %s", cb.State())
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, tc.run)
	}
}
