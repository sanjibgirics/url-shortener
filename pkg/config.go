package pkg

// RequestURL structure
type RequestURL struct {
	URL string `json:"url"`
}

// ResponseURL structure
type ResponseURL struct {
	OriginalURL string `json:"originalURL"`
	ShortURL    string `json:"shortURL"`
}
