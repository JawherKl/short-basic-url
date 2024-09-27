package main

import (
    "net/http"
)

// Handle redirecting the shortened URL
func handleRedirect(w http.ResponseWriter, r *http.Request) {
    // Get the short URL ID from the URL path
    shortID := r.URL.Path[1:]

    // Look up the original URL in the store
    originalURL, ok := urlStore[shortID]
    if !ok {
        http.Error(w, "Short URL not found", http.StatusNotFound)
        return
    }

    // Redirect to the original URL
    http.Redirect(w, r, originalURL, http.StatusFound)
}
