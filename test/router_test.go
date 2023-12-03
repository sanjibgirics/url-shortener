package pkg

import (
	"bytes"
	"encoding/json"
	"github.com/sanjibgirics/url-shortener/pkg"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestHomePage(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8083/", nil)
	assert.Nil(t, err)

	c := http.Client{}
	resp, err := c.Do(req)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	// Verify the response data
	responseBody, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, string(responseBody), "Welcome to URL Shortener !!!")
}

// TestUrlShortener verify major functionalities like shortening and getting back original
// url is working fine
func TestUrlShortener(t *testing.T) {
	// Call shortening url endpoint

	// Create payload
	originalURL := "http://example.com"
	requestURL := &pkg.RequestURL{URL: originalURL}

	requestURLJson, err := json.Marshal(requestURL)
	assert.Nil(t, err)
	assert.NotNil(t, requestURLJson)

	// Create HTTP request
	requestBody := bytes.NewBuffer(requestURLJson)
	req, err := http.NewRequest("POST", "http://localhost:8083/shorturl", requestBody)
	assert.Nil(t, err)
	assert.NotNil(t, req)
	req.Header.Set("Content-Type", "application/json")

	// Create client
	c := http.Client{}
	resp, err := c.Do(req)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Body)

	// Extract short url and long url
	responseURL := &pkg.ResponseURL{}
	err = json.NewDecoder(resp.Body).Decode(responseURL)
	assert.Nil(t, err)
	shortURL := responseURL.ShortURL
	assert.Equal(t, responseURL.OriginalURL, originalURL)
	assert.NotEqual(t, shortURL, "")

	// Call getting back original url endpoint

	requestURL.URL = shortURL

	requestURLJson, err = json.Marshal(requestURL)
	assert.Nil(t, err)
	assert.NotNil(t, requestURLJson)

	// Create HTTP request
	requestBody = bytes.NewBuffer(requestURLJson)
	req, err = http.NewRequest("POST", "http://localhost:8083/originalurl", requestBody)
	assert.Nil(t, err)
	assert.NotNil(t, req)
	req.Header.Set("Content-Type", "application/json")

	resp, err = c.Do(req)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Body)

	// Extract short url and long url
	err = json.NewDecoder(resp.Body).Decode(responseURL)
	assert.Nil(t, err)
	assert.Equal(t, responseURL.ShortURL, shortURL)

	// Ultimately verify that the url we got back using the short url is same as
	// original url which we had shortened
	assert.Equal(t, responseURL.OriginalURL, originalURL)
}
