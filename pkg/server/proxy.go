package server

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

func handleProxy(w http.ResponseWriter, r *http.Request) {
	// Example: Fetching a response from an external URL.
	resp, err := http.Get("http://example.com")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy response headers (optional)
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Set the status code from the fetched response.
	w.WriteHeader(resp.StatusCode)

	// Stream the response body to the client.
	if _, err := io.Copy(w, resp.Body); err != nil {
		slog.Error(fmt.Sprintf("Error while copying response: %v", err))
	}
}
