package files

import "strings"

// DataProcessor is the Abstract class, it is used in the ProcessContent() method.
type DataProcessor interface {
	Process() string
}

// RequestProcessor is the Abstract class,
// the ProcessContent is the only specific algorithm that needs to implement for different concrete classes.
// the param in ProcessContent is the interface MessageRetriever, it is used as an adapter,
// MessageRetriever can be used to convert data, and don't need to update the interface.
type RequestProcessor interface {
	getFirstPart() string
	getLastPart() string
	ProcessContent(DataProcessor) string
}

// WordProcessor struct is as the Concrete class 1, methods of RequestProcessor are implemented under below.
type WordProcessor struct{}

func (w *WordProcessor) getFirstPart() string {
	return "hello"
}

func (w *WordProcessor) getLastPart() string {
	return "template"
}

func (w *WordProcessor) ProcessContent(m DataProcessor) string {
	return strings.Join([]string{w.getFirstPart(), m.Process(), w.getLastPart()}, " ")
}

type PdfProcessor struct{}

func (p *PdfProcessor) getFirstPart() string {
	return "hello"
}

func (p *PdfProcessor) getLastPart() string {
	return "template"
}

func (p *PdfProcessor) ProcessContent(f func() string) string {
	return strings.Join([]string{p.getFirstPart(), f(), p.getLastPart()}, " ")
}

type dataAdapter struct {
	myFunc func() string
}

func (d *dataAdapter) Process() string {
	if d.myFunc != nil {
		return d.myFunc()
	}

	return ""
}

func DataProcessorAdapter(f func() string) DataProcessor {
	return &dataAdapter{myFunc: f}
}
