package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGenerateShortURL(t *testing.T) {
	originalURL := "https://github.com/rajjmalla10"

	// Call the function to test
	shortURL := generateShortURL(originalURL)

	// Assert that the generated shortURL is not empty
	if shortURL == "" {
		t.Errorf("generateShortURL() = %s, want not empty", shortURL)
	}
}

func TestShortURLHandler(t *testing.T) {
	// Create a request to pass to our handler
	requestBody := `{"url": "https://github.com/rajjmalla10"}`
	req, err := http.NewRequest("POST", "/shorten", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ShortURLHandler)

	// Serve the HTTP request to the ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode the response body into a struct
	var response struct {
		ShortURL string `json:"short_url"`
	}
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("failed to decode JSON response body: %v", err)
	}

	// Generate the expected short URL
	expectedShortURL := generateShortURL("https://github.com/rajjmalla10")

	// Compare the generated short URL with the response
	if response.ShortURL != expectedShortURL {
		t.Errorf("handler returned unexpected short URL: got %v want %v",
			response.ShortURL, expectedShortURL)
	}
}
