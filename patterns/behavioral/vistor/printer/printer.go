package printer

import (
	"fmt"
	"io"
	"os"
)

// Printer is the abstract interface
type Printer interface {
	Print()
}

// Visitor is the abstract interface, all visit methods are defined in it.
// Methods in Visitor interface can be used to alter the Visitable object.
// Using this way, Visitor can change the original functionalities indirectly, and it won't change the original interface source code.
type Visitor interface {
	VisitA(*MessageA)
	VisitB(*MessageB)
}

// Visitable is the abstract interface. The key point of the visitor pattern, Accept(Visitor) is defined in it.
// It can accept many visitors and does not need to change the original interface, the original interface at here is the Printer interface.
type Visitable interface {
	Accept(Visitor)
}

// MessageA is a concrete struct
type MessageA struct {
	Msg    string
	Output io.Writer
}

// Print method help us to test the types, it must print to the Stdout call by default.
func (m *MessageA) Print() {
	// It first checks the content of the Output field to assign the output of the os.Stdout call in case it is null.
	// In tests, we are storing a pointer there to our TestHelper type so this line is never executed in our test.
	if m.Output == nil {
		m.Output = os.Stdout
	}

	// Using Fprintf method takes an io.Writer package as the first argument and the text to format as the next arguments.
	// Finally, each message type prints to the Output field, the full message stored in the Msg field.
	fmt.Fprintf(m.Output, "A: %s", m.Msg)
}

func (m *MessageA) Accept(v Visitor) {
	v.VisitA(m)
}

// MessageB is a concrete struct
type MessageB struct {
	Msg    string
	Output io.Writer
}

func (m *MessageB) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}

	fmt.Fprintf(m.Output, "B: %s", m.Msg)
}

func (m *MessageB) Accept(v Visitor) {
	v.VisitB(m)
}

type MessageVisitor struct{}

// VisitA uses the fmt.Sprintf method returns a formatted string with the actual contents of m.Msg, a white space, and the message, Visited
// This string will be stored on the Msg field, overriding the previous contents.
func (mv *MessageVisitor) VisitA(m *MessageA) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited A)")
}

// VisitB uses the fmt.Sprintf method returns a formatted string with the actual contents of m.Msg, a white space, and the message, Visited
// This string will be stored on the Msg field, overriding the previous contents.
func (mv *MessageVisitor) VisitB(m *MessageB) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited B)")
}

type MsgFieldVisitorPrinter struct{}

func (mf *MsgFieldVisitorPrinter) VisitA(m *MessageA) {
	fmt.Printf(m.Msg)
}
func (mf *MsgFieldVisitorPrinter) VisitB(m *MessageB) {
	fmt.Printf(m.Msg)
}
