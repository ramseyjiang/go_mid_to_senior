package microcb

import (
	"time"

	"github.com/sony/gobreaker"
)

// ExternalServiceFunc represents a function that simulates an external service call.
type ExternalServiceFunc func() (string, error)

// CircuitBreakerConfig holds configuration for the circuit breaker.
type CircuitBreakerConfig struct {
	Timeout          time.Duration
	MaxRequests      uint32
	Interval         time.Duration
	FailureThreshold uint32
}

// ServiceClient provides an interface to call the external service with a circuit breaker.
type ServiceClient struct {
	cb *gobreaker.CircuitBreaker
}

// NewServiceClient creates a new ServiceClient with the given circuit breaker configuration.
func NewServiceClient(cfg CircuitBreakerConfig) *ServiceClient {
	settings := gobreaker.Settings{
		Name:        "ServiceClient",
		MaxRequests: cfg.MaxRequests,
		Interval:    cfg.Interval,
		Timeout:     cfg.Timeout,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= cfg.FailureThreshold && failureRatio >= 0.6
		},
	}
	cb := gobreaker.NewCircuitBreaker(settings)
	return &ServiceClient{cb: cb}
}

// CallService uses the circuit breaker to call the external service.
func (c *ServiceClient) CallService(service ExternalServiceFunc) (string, error) {
	result, err := c.cb.Execute(func() (interface{}, error) {
		return service()
	})

	if err != nil {
		return "", err
	}
	return result.(string), nil
}
