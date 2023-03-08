package sample

// MixinInterface is the mixin interface with the methods that you want to reuse.
type MixinInterface interface {
	MixinMethod() string
}

// Mixin is the mixin struct  that implements the mixin interface.
type Mixin struct{}

func (m *Mixin) MixinMethod() string {
	return "This is the Mixin Method"
}

// Receiver is the receiver struct that will use the behavior of the mixin.
type Receiver struct {
	MixinInterface
}

func (r *Receiver) ReceiverMethod() string {
	return "This is the Receiver Method"
}

// NewReceiver is a function created to initialize the receiver struct with the mixin struct.
func NewReceiver() *Receiver {
	return &Receiver{&Mixin{}}
}
