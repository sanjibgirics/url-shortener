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

// DomainUsage structure contains domain name and number time it has used the shortener service
type DomainUsage struct {
	Domain string
	Usage  int
}

// TopDomainMetrics structure
type TopDomainMetrics []DomainUsage
