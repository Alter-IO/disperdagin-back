package util

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateRandomPassword creates a random 12-character password
func GenerateRandomPassword() (string, error) {
	// Generate 9 random bytes which will become 12 base64 characters
	randomBytes := make([]byte, 9)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Convert to base64 (which will be ~12 chars)
	password := base64.StdEncoding.EncodeToString(randomBytes)

	// Take only the first 12 chars to ensure consistent length
	if len(password) > 12 {
		password = password[:12]
	}

	return password, nil
}
