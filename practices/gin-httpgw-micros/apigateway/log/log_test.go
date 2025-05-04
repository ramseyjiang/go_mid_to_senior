package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReceiveLog(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Valid Log",
			input:    `{"action": "User added", "user": {"id": "123", "name": "Test User"}}`,
			expected: http.StatusOK,
		},
		{
			name:     "Empty Log",
			input:    "",
			expected: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/log", bytes.NewBufferString(test.input))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(receiveLog)
			handler.ServeHTTP(rr, req)

			assert.Equal(t, test.expected, rr.Code)
		})
	}
}
