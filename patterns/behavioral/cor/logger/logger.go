package logger

import (
	"log"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	ERROR
)

type Handler interface {
	SetNext(Handler)
	HandleLog(level LogLevel, message string)
}

type DebugHandler struct {
	next   Handler
	logger *log.Logger
}

func (h *DebugHandler) SetNext(next Handler) {
	h.next = next
}

func (h *DebugHandler) HandleLog(level LogLevel, message string) {
	if level == DEBUG {
		h.logger.Println("[DEBUG]:", message)
	}

	if h.next != nil {
		h.next.HandleLog(level, message)
	}
}

type InfoHandler struct {
	next   Handler
	logger *log.Logger
}

func (h *InfoHandler) SetNext(next Handler) {
	h.next = next
}

func (h *InfoHandler) HandleLog(level LogLevel, message string) {
	if level == INFO {
		h.logger.Println("[INFO]:", message)
	}

	if h.next != nil {
		h.next.HandleLog(level, message)
	}
}

type ErrorHandler struct {
	next   Handler
	logger *log.Logger
}

func (h *ErrorHandler) SetNext(next Handler) {
	h.next = next
}

func (h *ErrorHandler) HandleLog(level LogLevel, message string) {
	if level == ERROR {
		h.logger.Println("[ERROR]:", message)
	}

	if h.next != nil {
		h.next.HandleLog(level, message)
	}
}

func createHandlerChain(debugLogger, infoLogger, errorLogger *log.Logger) Handler {
	debugHandler := &DebugHandler{logger: debugLogger}
	infoHandler := &InfoHandler{logger: infoLogger}
	errorHandler := &ErrorHandler{logger: errorLogger}

	debugHandler.SetNext(infoHandler)
	infoHandler.SetNext(errorHandler)

	return debugHandler
}
