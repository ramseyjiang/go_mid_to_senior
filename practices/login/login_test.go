package login

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name string
		pass string
	}{
		{name: "basic password", pass: "password123"},
		{name: "empty password", pass: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashed := hashPassword(tt.pass)
			if hashed == tt.pass {
				t.Errorf("hashPassword() returns plain text")
			}

			// Separate check for consistent hashing
			if hashPassword(tt.pass) != hashPassword(tt.pass) {
				t.Errorf("hashPassword() produces inconsistent results")
			}
		})
	}
}

func TestValidateCredentials(t *testing.T) {
	tests := []struct {
		name     string
		username string
		password string
		want     bool
	}{
		{
			name:     "valid credentials",
			username: "testuser",
			password: "correctpassword",
			want:     true,
		},
		{
			name:     "wrong password",
			username: "testuser",
			password: "wrongpassword",
			want:     false,
		},
		{
			name:     "nonexistent user",
			username: "nonexistentuser",
			password: "password",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateCredentials(tt.username, tt.password); got != tt.want {
				t.Errorf("ValidateCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}
