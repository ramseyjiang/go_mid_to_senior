package logger

// func TestChain(t *testing.T) {
// 	handlerChain := createHandlerChain()
//
// 	handlerChain.HandleLog(DEBUG, "This is a debug message")
// 	handlerChain.HandleLog(INFO, "This is an info message")
// 	handlerChain.HandleLog(ERROR, "This is an error message")
// }

import (
	"bytes"
	"log"
	"testing"
)

func TestHandlerChain(t *testing.T) {
	tests := []struct {
		name    string
		level   LogLevel
		message string
		want    string
	}{
		{"DebugMessage", DEBUG, "Debug Test", "[DEBUG]: Debug Test\n"},
		{"InfoMessage", INFO, "Info Test", "[INFO]: Info Test\n"},
		{"ErrorMessage", ERROR, "Error Test", "[ERROR]: Error Test\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			logger := log.New(&buf, "", 0)

			handlerChain := createHandlerChain(logger, logger, logger)
			handlerChain.HandleLog(tt.level, tt.message)

			got := buf.String()
			if got != tt.want {
				t.Errorf("HandleLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
