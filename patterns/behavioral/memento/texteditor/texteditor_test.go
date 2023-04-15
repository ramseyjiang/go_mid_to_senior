package texteditor

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestUndoRedoManager(t *testing.T) {
	textEditor := &TextEditor{}
	undoRedoManager := NewUndoRedoManager()

	textEditor.SetText("Hello")
	undoRedoManager.SaveState(textEditor)

	textEditor.SetText("Hello, world!")
	undoRedoManager.SaveState(textEditor)

	textEditor.SetText("Hello, Golang!")
	undoRedoManager.SaveState(textEditor)

	assert.Equal(t, "Hello, Golang!", textEditor.GetText())
	if textEditor.GetText() != "Hello, Golang!" {
		t.Errorf("Expected text to be 'Hello, Golang!', but got '%s'", textEditor.GetText())
	}

	undoRedoManager.Undo(textEditor)
	assert.Equal(t, "Hello, world!", textEditor.GetText())
	if textEditor.GetText() != "Hello, world!" {
		t.Errorf("Expected text to be 'Hello, world!', but got '%s'", textEditor.GetText())
	}

	undoRedoManager.Undo(textEditor)
	assert.Equal(t, "Hello", textEditor.GetText())
	if textEditor.GetText() != "Hello" {
		t.Errorf("Expected text to be 'Hello', but got '%s'", textEditor.GetText())
	}

	undoRedoManager.Redo(textEditor)
	assert.Equal(t, "Hello, world!", textEditor.GetText())
	if textEditor.GetText() != "Hello, world!" {
		t.Errorf("Expected text to be 'Hello, world!', but got '%s'", textEditor.GetText())
	}

	undoRedoManager.Redo(textEditor)
	assert.Equal(t, "Hello, Golang!", textEditor.GetText())
	if textEditor.GetText() != "Hello, Golang!" {
		t.Errorf("Expected text to be 'Hello, Golang!', but got '%s'", textEditor.GetText())
	}

	// Test redo boundary
	undoRedoManager.Redo(textEditor)
	assert.Equal(t, "Hello, Golang!", textEditor.GetText())
	if textEditor.GetText() != "Hello, Golang!" {
		t.Errorf("Expected text to be 'Hello, Golang!', but got '%s'", textEditor.GetText())
	}

	// Test undo boundary
	undoRedoManager.Undo(textEditor)
	assert.Equal(t, "Hello, world!", textEditor.GetText())
	undoRedoManager.Undo(textEditor)
	assert.Equal(t, "Hello", textEditor.GetText())
	undoRedoManager.Undo(textEditor)
	assert.Equal(t, "Hello", textEditor.GetText())
	if textEditor.GetText() != "Hello" {
		t.Errorf("Expected text to be 'Hello', but got '%s'", textEditor.GetText())
	}
}
