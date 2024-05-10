package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateCustomRelativePath generates a random base64-encoded string of a specified length.
// The function calculates the number of random bytes needed to achieve the desired length,
// generates random bytes, encodes them to base64, and trims the result to the desired length.
func GenerateCustomRelativePath(pathLength int64) (string, error) {

	// Calculate how many random bytes we need to generate to achieve the desired length.
	numBytes := (pathLength * 6) / 8 // Each base64 character represents 6 bits.

	// Generate random bytes.
	bytes := make([]byte, numBytes)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to a base64 string.
	relativePath := base64.URLEncoding.EncodeToString(bytes)

	// Trim the string to the desired length.
	if int64(len(relativePath)) > pathLength {
		relativePath = relativePath[:pathLength]
	}

	return relativePath, nil
}
