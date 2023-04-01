package principle

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestHandlerA(t *testing.T) {
	handler := createHandlerChain()

	// create a Request structure and pass it to the first handler in the chain (h1) using the Handle method.
	req := &Request{
		Value: 5,
	}
	assert.Equal(t, "ConcreteHandlerA handled the request", handler.Handle(req))
}

func TestHandlerB(t *testing.T) {
	handler := createHandlerChain()

	// create a Request structure and pass it to the first handler in the chain (h1) using the Handle method.
	req := &Request{
		Value: 15,
	}

	assert.Equal(t, "ConcreteHandlerB handled the request", handler.Handle(req))
}
