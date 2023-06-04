package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReceiveLog(t *testing.T) {
	logMessage := `{"action": "User added", "user": {"id": 1, "name": "Test User"}}`

	req, err := http.NewRequest("POST", "/log", bytes.NewBuffer([]byte(logMessage)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(receiveLog)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
