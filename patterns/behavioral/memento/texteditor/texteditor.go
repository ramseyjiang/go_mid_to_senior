package texteditor

type TextEditor struct {
	text string
}

func (te *TextEditor) SetText(text string) {
	te.text = text
}

func (te *TextEditor) GetText() string {
	return te.text
}

func (te *TextEditor) CreateMemento() *Memento {
	return &Memento{text: te.text}
}

func (te *TextEditor) Restore(m *Memento) {
	te.text = m.GetText()
}

type Memento struct {
	text string
}

func (m *Memento) GetText() string {
	return m.text
}

type History struct {
	mementos []*Memento
}

func (h *History) AddMemento(m *Memento) {
	h.mementos = append(h.mementos, m)
}

func (h *History) GetMemento(index int) *Memento {
	if index < 0 || index >= len(h.mementos) {
		return nil
	}
	return h.mementos[index]
}

type UndoRedoManager struct {
	history    *History
	currentIdx int
}

func NewUndoRedoManager() *UndoRedoManager {
	history := &History{}
	return &UndoRedoManager{
		history:    history,
		currentIdx: -1,
	}
}

func (urm *UndoRedoManager) SaveState(te *TextEditor) {
	urm.history.AddMemento(te.CreateMemento())
	urm.currentIdx++
}

func (urm *UndoRedoManager) Undo(te *TextEditor) {
	if urm.currentIdx <= 0 {
		return
	}
	urm.currentIdx--
	te.Restore(urm.history.GetMemento(urm.currentIdx))
}

func (urm *UndoRedoManager) Redo(te *TextEditor) {
	if urm.currentIdx >= len(urm.history.mementos)-1 {
		return
	}
	urm.currentIdx++
	te.Restore(urm.history.GetMemento(urm.currentIdx))
}
