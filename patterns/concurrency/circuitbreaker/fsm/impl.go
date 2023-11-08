package fsm

import (
	"errors"
	"sync"
	"time"
)

// State represents the state of the CircuitBreaker.
type State int

const (
	// Closed state allows calls to pass through.
	Closed State = iota
	// Open state blocks all calls.
	Open
	// HalfOpen state allows a limited number of calls to test the health of the service.
	HalfOpen
)

// String returns the string representation of the State.
func (s State) String() string {
	return [...]string{"Closed", "Open", "HalfOpen"}[s]
}

// CircuitBreaker struct holds the state of the circuit breaker.
type CircuitBreaker struct {
	state            State
	failureCount     int
	failureThreshold int
	resetTimeout     time.Duration
	mu               sync.Mutex
}

// NewCircuitBreaker creates a new CircuitBreaker with the given failure threshold and reset timeout.
func NewCircuitBreaker(failureThreshold int, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:            Closed,
		failureThreshold: failureThreshold,
		resetTimeout:     resetTimeout,
	}
}

// State returns the current state of the CircuitBreaker.
func (cb *CircuitBreaker) State() State {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.state
}

// Success resets the failure count when a call succeeds.
func (cb *CircuitBreaker) Success() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	if cb.state == HalfOpen {
		cb.state = Closed
	}
	cb.failureCount = 0
}

// Failure increments the failure count and transitions the state if the threshold is reached.
func (cb *CircuitBreaker) Failure() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.failureCount++
	if cb.failureCount >= cb.failureThreshold {
		cb.transitionToOpen()
	}
}

// transitionToOpen sets the state to Open and starts the reset timeout.
func (cb *CircuitBreaker) transitionToOpen() {
	cb.state = Open
	go cb.startResetTimeout()
}

// startResetTimeout transitions the state to HalfOpen after the reset timeout.
func (cb *CircuitBreaker) startResetTimeout() {
	time.Sleep(cb.resetTimeout)
	cb.mu.Lock()
	defer cb.mu.Unlock()
	if cb.state == Open {
		cb.state = HalfOpen
	}
}

// AttemptExecution checks if a call should be allowed to proceed based on the current state.
func (cb *CircuitBreaker) AttemptExecution() error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	switch cb.state {
	case Open:
		return errors.New("circuit breaker is open")
	case HalfOpen:
		// Allow only one call to go through when HalfOpen.
		cb.state = Open
		return nil
	case Closed:
		return nil
	default:
		return errors.New("invalid circuit breaker state")
	}
}
