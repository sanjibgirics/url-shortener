package pkg

import (
	"strings"
)

const baseURL = "https://myurl.io/"

// processShorteningURL processes the shortening of the original url
func processShorteningURL(originalURL string) *ResponseURL {
	responseURL := &ResponseURL{}

	originalURL = strings.Trim(originalURL, " ")
	// Check if the original URL's short version already present.
	// If so, return the short url directly respective to that.
	if _, ok := originalURLToShortCode[originalURL]; ok {
		responseURL.OriginalURL = originalURL
		responseURL.ShortURL = baseURL + originalURLToShortCode[originalURL]
	} else {
		shortCode := generateShortCode()
		shortURL := baseURL + shortCode

		responseURL.OriginalURL = originalURL
		responseURL.ShortURL = shortURL

		// store the data
		shortCodeToOriginalURL[shortCode] = originalURL
		originalURLToShortCode[originalURL] = shortCode
	}

	return responseURL
}

// processGetOriginalURL returns the original url from the given short url
func processGetOriginalURL(shortURL string) *ResponseURL {
	responseURL := &ResponseURL{}

	shortURLTrimmed := strings.Trim(shortURL, " ")
	// Check baseURL is present and in proper place
	if shortURLTrimmed[:len(baseURL)] != baseURL {
		return nil
	}
	shortCode := shortURL[len(baseURL):]

	if _, ok := shortCodeToOriginalURL[shortCode]; ok {
		responseURL.OriginalURL = shortCodeToOriginalURL[shortCode]
		responseURL.ShortURL = shortURL
	} else {
		return nil
	}

	return responseURL
}
