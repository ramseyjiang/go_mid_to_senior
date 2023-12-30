package microcb

import (
	"errors"
	"testing"
	"time"
)

func TestServiceClient_CircuitBreakerStates(t *testing.T) {
	failingService := func() (string, error) {
		return "", errors.New("service failure")
	}

	healthyService := func() (string, error) {
		return "success", nil
	}

	// Increase the delay to ensure it exceeds the circuit breaker's timeout
	slowService := func() (string, error) {
		time.Sleep(200 * time.Millisecond) // Deliberately slow response
		return "slow response", errors.New("service timeout")
	}

	cfg := CircuitBreakerConfig{
		Timeout:          50 * time.Millisecond,
		MaxRequests:      1,
		Interval:         500 * time.Millisecond,
		FailureThreshold: 3,
	}

	tests := []struct {
		name            string
		serviceFunc     ExternalServiceFunc
		triggerFailures int
		waitDuration    time.Duration
		expectError     bool
	}{
		{
			name:            "Open State - After Failures",
			serviceFunc:     failingService,
			triggerFailures: int(cfg.FailureThreshold),
			waitDuration:    0,
			expectError:     true,
		},
		{
			name:            "Half-Open State - Allows Trial Request",
			serviceFunc:     healthyService,
			triggerFailures: int(cfg.FailureThreshold),
			waitDuration:    cfg.Timeout + 10*time.Millisecond,
			expectError:     false,
		},
		{
			name:            "Closed State - After Successful Request",
			serviceFunc:     healthyService,
			triggerFailures: 0,
			waitDuration:    0,
			expectError:     false,
		},
		{
			name:            "Slow Service",
			serviceFunc:     slowService,
			triggerFailures: 0,
			waitDuration:    0,
			expectError:     true, // Expect an error due to timeout
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewServiceClient(cfg)

			// Trigger failures if necessary
			for i := 0; i < tt.triggerFailures; i++ {
				client.CallService(failingService)
			}

			// Wait if necessary
			if tt.waitDuration > 0 {
				time.Sleep(tt.waitDuration)
			}

			// Call the service and check for expected error
			_, err := client.CallService(tt.serviceFunc)
			if (err != nil) != tt.expectError {
				t.Errorf("Test '%s' failed: expected error = %v, got error = %v", tt.name, tt.expectError, err != nil)
			}
		})
	}
}
