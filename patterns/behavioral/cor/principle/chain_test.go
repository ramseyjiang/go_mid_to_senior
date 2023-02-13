package principle

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestHandlerA(t *testing.T) {
	h1 := &ConcreteHandlerA{}
	h2 := &ConcreteHandlerB{}
	h1.SetNext(h2)

	// create a Request structure and pass it to the first handler in the chain (h1) using the Handle method.
	req := &Request{
		Value: 5,
	}
	assert.Equal(t, "ConcreteHandlerA handled the request", h1.Handle(req))
}

func TestHandlerB(t *testing.T) {
	h1 := &ConcreteHandlerA{}
	h2 := &ConcreteHandlerB{}
	h1.SetNext(h2)

	// create a Request structure and pass it to the first handler in the chain (h1) using the Handle method.
	req := &Request{
		Value: 15,
	}

	assert.Equal(t, "ConcreteHandlerB handled the request", h1.Handle(req))
}
