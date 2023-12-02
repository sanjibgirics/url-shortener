package pkg

import (
	"math/rand"
	"net/url"
	"time"
)

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const codeLength = 6

// isValidURL checks whether given url is valid or not.
func isValidURL(urlString string) bool {
	parsedURL, err := url.Parse(urlString)
	if err != nil || parsedURL.Host == "" {
		return false
	}
	return true
}

// generateShortCode generates a random short code
func generateShortCode() string {
	rand.NewSource(time.Now().UnixNano())
	shortCode := make([]byte, codeLength)
	for i := range shortCode {
		shortCode[i] = characters[rand.Intn(len(characters))]
	}

	return string(shortCode)
}
