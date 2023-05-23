package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func main() {
	// Generate a random token
	token, err := generateRandomToken(32)
	if err != nil {
		log.Fatalf("Failed to generate random token: %s", err)
	}

	log.Printf("Generated token: %s", token)
}

// generateRandomToken generates a random token with the specified length
func generateRandomToken(length int) (string, error) {
	// Calculate the number of bytes needed
	numBytes := (length * 6) / 8
	if (length*6)%8 != 0 {
		numBytes++
	}

	// Generate random bytes
	bytes := make([]byte, numBytes)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Convert bytes to base64 string
	token := base64.URLEncoding.EncodeToString(bytes)

	// Trim the string to the desired length
	if len(token) > length {
		token = token[:length]
	}

	return token, nil
}
