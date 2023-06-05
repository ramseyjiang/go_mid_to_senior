package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	tests := []struct {
		name     string
		input    User
		expected int
	}{
		{
			name:     "Valid User",
			input:    User{Name: "Test User"},
			expected: http.StatusOK,
		},
		{
			name:     "Empty User",
			input:    User{},
			expected: http.StatusBadRequest,
		},
		{
			name:     "Empty Username",
			input:    User{Name: ""},
			expected: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reqBody, _ := json.Marshal(test.input)
			req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(addUser)
			handler.ServeHTTP(rr, req)

			assert.Equal(t, test.expected, rr.Code)
		})
	}
}

func TestGetUsers(t *testing.T) {
	tests := []struct {
		name     string
		users    []User
		expected int
	}{
		{
			name:     "No Users",
			users:    []User{},
			expected: http.StatusOK,
		},
		{
			name: "With Users",
			users: []User{
				{ID: uuid.New().String(), Name: "Test User 1"},
				{ID: uuid.New().String(), Name: "Test User 2"},
			},
			expected: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			users = test.users

			req, err := http.NewRequest("GET", "/user", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(getUsers)
			handler.ServeHTTP(rr, req)

			assert.Equal(t, test.expected, rr.Code)
		})
	}
}
