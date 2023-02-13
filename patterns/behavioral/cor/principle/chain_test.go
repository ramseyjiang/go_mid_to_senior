package principle

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestHandle(t *testing.T) {
	h1 := &ConcreteHandlerA{}
	h2 := &ConcreteHandlerB{}
	h1.SetNext(h2)

	// create a Request structure and pass it to the first handler in the chain (h1) using the Handle method.
	req := &Request{
		Type:  "A",
		Value: 100,
	}
	assert.Equal(t, "ConcreteHandlerA handled the request", h1.Handle(req))
}
