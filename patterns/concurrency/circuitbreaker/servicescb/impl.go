package servicescb

import (
	"errors"
	"sync"
	"time"
)

// ServiceFunc is a function that represents a service call.
type ServiceFunc func() error

// CircuitBreaker is a custom circuit breaker.
type CircuitBreaker struct {
	failureThreshold int
	recoveryTimeout  time.Duration
	lock             sync.Mutex
	failures         int
	state            string
	lastFailureTime  time.Time
}

// NewCircuitBreaker creates a new CircuitBreaker.
func NewCircuitBreaker(failureThreshold int, recoveryTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		failureThreshold: failureThreshold,
		recoveryTimeout:  recoveryTimeout,
		state:            "closed",
	}
}

// Call wraps a service call with the circuit breaker logic.
func (cb *CircuitBreaker) Call(service ServiceFunc) error {
	cb.lock.Lock()
	defer cb.lock.Unlock()

	if cb.state == "open" {
		if time.Since(cb.lastFailureTime) > cb.recoveryTimeout {
			cb.state = "half-open"
		} else {
			return errors.New("circuit breaker open")
		}
	}

	err := service()
	if err != nil {
		cb.failures++
		if cb.failures >= cb.failureThreshold {
			cb.state = "open"
			cb.lastFailureTime = time.Now()
		}
		return err
	}

	cb.reset()
	return nil
}

func (cb *CircuitBreaker) reset() {
	cb.failures = 0
	cb.state = "closed"
}

// ServiceA simulates a service call.
func ServiceA() error {
	// Simulate service logic
	return nil // or return an error to simulate failure
}

// ServiceB simulates another service call.
func ServiceB() error {
	// Simulate service logic
	return nil // or return an error to simulate failure
}
