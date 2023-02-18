package printer

import "testing"

// The TestHelper struct implements the io.Writer interface. It stores the written bytes on the Received field
// It is used to check the contents of Received to test against our expected value.
type TestHelper struct {
	Received string
}

func (t *TestHelper) Write(p []byte) (int, error) {
	t.Received = string(p)
	return len(p), nil
}

func TestAll(t *testing.T) {
	testHelper := &TestHelper{}
	visitor := &MessageVisitor{}

	t.Run("MessageA test", func(t *testing.T) {
		msg := MessageA{
			Msg:    "Hello World",
			Output: testHelper,
		}

		// Inside to Accept(Visitor) method on the MessageA struct,
		// the VisitA(*MessageA) method is executed to alter the contents of the Msg field.
		// That's why we passed the pointer to VisitA method, without a pointer the contents won't be persisted.
		msg.Accept(visitor)

		// To test if the Visitor type has done its job within the Accept method, we must call the Print() method on the MessageA type.
		// Using the Print() method, the MessageA struct must write the contents of Msg to the provided io.Writer interface.
		msg.Print()

		expected := "A: Hello World (Visited A)"
		if testHelper.Received != expected {
			t.Errorf("Expected result was incorrect. %s != %s",
				testHelper.Received, expected)
		}
	})

	t.Run("MessageB test", func(t *testing.T) {
		msg := MessageB{
			Msg:    "Hello World",
			Output: testHelper,
		}

		msg.Accept(visitor)
		msg.Print()

		expected := "B: Hello World (Visited B)"
		if testHelper.Received != expected {
			t.Errorf("Expected result was incorrect. %s != %s",
				testHelper.Received, expected)
		}
	})

	msgFieldVisitor := &MsgFieldVisitorPrinter{}
	t.Run("MessageA MsgFieldVisitorPrinter test", func(t *testing.T) {
		msg := MessageA{
			Msg:    "Hello World",
			Output: testHelper,
		}

		msg.Accept(msgFieldVisitor)
		msg.Print()

		expected := "A: Hello World"
		if testHelper.Received != expected {
			t.Errorf("Expected result was incorrect. %s != %s",
				testHelper.Received, expected)
		}
	})

	t.Run("MessageB MsgFieldVisitorPrinter test", func(t *testing.T) {
		msg := MessageB{
			Msg:    "Hello World",
			Output: testHelper,
		}

		msg.Accept(msgFieldVisitor)
		msg.Print()

		expected := "B: Hello World"
		if testHelper.Received != expected {
			t.Errorf("Expected result was incorrect. %s != %s",
				testHelper.Received, expected)
		}
	})
}
