package pkg

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// RegisterRoutes register all the endpoints needed  to the router
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/shorturl", shortURLHandler).Methods("POST")
	r.HandleFunc("/originalurl", originalURLHandler).Methods("POST")
}

// shortURL will return the shortened url version of original url
func shortURLHandler(w http.ResponseWriter, r *http.Request) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		http.Error(w, "unsupported media type", http.StatusUnsupportedMediaType)
		return
	}

	// Retrieve original url

	requestURL := &RequestURL{}
	if err := json.NewDecoder(r.Body).Decode(requestURL); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if requestURL.URL == "" {
		http.Error(w, "url not given", http.StatusBadRequest)
		return
	}

	if !isValidURL(requestURL.URL) {
		http.Error(w, "url given is not valid", http.StatusBadRequest)
		return
	}

	// Shorten the url
	responseURL := processShorteningURL(requestURL.URL)

	// Prepare response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responseURL); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// originalURLHandler will return the original url respective to the short url
func originalURLHandler(w http.ResponseWriter, r *http.Request) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		http.Error(w, "unsupported media type", http.StatusUnsupportedMediaType)
		return
	}

	// Retrieve short url

	requestURL := &RequestURL{}
	if err := json.NewDecoder(r.Body).Decode(requestURL); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if requestURL.URL == "" {
		http.Error(w, "url not given", http.StatusBadRequest)
		return
	}

	if !isValidURL(requestURL.URL) {
		http.Error(w, "url given is not valid", http.StatusBadRequest)
		return
	}

	// Getting back the original url
	responseURL := processGetOriginalURL(requestURL.URL)
	if responseURL == nil {
		http.Error(w, "given short url not found", http.StatusNotFound)
		return
	}

	// Prepare response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responseURL); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Handle homepage
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to URL Shortener !!!"))
}
