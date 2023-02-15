package files

import (
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

type TestStruct struct {
	WordProcessor
}

func (m *TestStruct) Process() string {
	return "world"
}

func TestTemplateProcess(t *testing.T) {
	t.Run("Using interfaces", func(t *testing.T) {
		s := &TestStruct{}
		res := s.ProcessContent(s)

		expectedOrError(res, " world ", t)
	})

	t.Run("Using anonymous function", func(t *testing.T) {
		m := new(PdfProcessor)
		res := m.ProcessContent(func() string { // anonymous function
			return "world"
		})

		expectedOrError(res, " world ", t)
	})

	t.Run("Using an anonymous function in adapter to adapt the interface", func(t *testing.T) {
		testDataAdapter := DataProcessorAdapter(func() string { // anonymous function
			return "world"
		})

		if testDataAdapter == nil {
			t.Fatal("Can not continue with a nil DataProcessor")
		}

		word := WordProcessor{}
		resWord := word.ProcessContent(testDataAdapter)

		pdf := PdfProcessor{}
		resPdf := pdf.ProcessContent(func() string { // anonymous function
			return "world"
		})

		expectedOrError(resWord, " world ", t)
		expectedOrError(resPdf, " world ", t)
	})
}

func expectedOrError(res string, expected string, t *testing.T) {
	if !strings.Contains(res, expected) {
		t.Errorf("Expected string '%s' was not found on returned string\n", expected)
	}
	assert.Equal(t, true, strings.Contains(res, expected))
}
