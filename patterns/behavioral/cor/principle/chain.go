package principle

type Request struct {
	Type  string
	Value int
}

// Handler interface defines the methods that all handlers must implement,
// including SetNext for setting the next handler in the chain, and Handle for handling the request.
type Handler interface {
	SetNext(Handler)
	Handle(*Request) string
}

// The ConcreteHandlerA and ConcreteHandlerB structures implement the Handler interface and represent the concrete handlers in the chain.
type ConcreteHandlerA struct {
	Next Handler
}

func (h *ConcreteHandlerA) SetNext(next Handler) {
	h.Next = next
}

// The Handle method of each handler checks whether it can handle the request
// and if not, passes the request to the next handler in the chain using the SetNext method.
func (h *ConcreteHandlerA) Handle(req *Request) (resp string) {
	if req.Value >= 0 && req.Value < 10 {
		return "ConcreteHandlerA handled the request"
	}
	if h.Next != nil {
		return h.Next.Handle(req)
	}
	return
}

type ConcreteHandlerB struct {
	Next Handler
}

func (h *ConcreteHandlerB) SetNext(next Handler) {
	h.Next = next
}

func (h *ConcreteHandlerB) Handle(req *Request) (resp string) {
	if req.Value >= 10 && req.Value < 20 {
		return "ConcreteHandlerB handled the request"
	}
	if h.Next != nil {
		return h.Next.Handle(req)
	}
	return
}

func createHandlerChain() Handler {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	handlerA.SetNext(handlerB)
	return handlerA
}
