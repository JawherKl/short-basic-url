package main

import (
    "encoding/json"
    "math/rand"
    "net/http"
    "time"
	"github.com/JawherKl/short-basic-url/utils"
)

// In-memory store for URLs
var urlStore = make(map[string]string)

// Handle shortening a URL
func handleShortenURL(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Parse the incoming URL
    var requestBody struct {
        URL string `json:"url"`
    }
    err := json.NewDecoder(r.Body).Decode(&requestBody)
    if err != nil || requestBody.URL == "" {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Generate a random ID for the short URL
    shortID := generateShortID()

    // Store the original URL with the short ID
    urlStore[shortID] = requestBody.URL

    // Respond with the shortened URL
    responseBody := struct {
        ShortURL string `json:"short_url"`
    }{
        ShortURL: "http://localhost:8080/" + shortID,
    }
    json.NewEncoder(w).Encode(responseBody)
}

// Generate a random string of length 6 for the short URL ID
func generateShortID() string {
    rand.Seed(time.Now().UnixNano())
    letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
    b := make([]rune, 6)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}
