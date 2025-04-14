package login

import (
	"crypto/sha256"
	"encoding/hex"
)

// User stores username and hashed password
type User struct {
	Username string
	Password string // Stored as a hashed password
}

// Mock user data
var users = map[string]User{
	"testuser": {Username: "testuser", Password: hashPassword("correctpassword")},
}

// HashPassword hashes the password using SHA-256
func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

// ValidateCredentials checks if the provided username and password match stored data
func ValidateCredentials(username, password string) bool {
	user, exists := users[username]
	if !exists {
		return false
	}
	return user.Password == hashPassword(password)
}
