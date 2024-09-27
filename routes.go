package main

import (
    "net/http"
    "github.com/gorilla/mux"
	"github.com/JawherKl/short-basic-url/handler"
	"github.com/JawherKl/short-basic-url/urlstore"
    
)

func routes() {
    http.HandleFunc("/shorten", handleShortenURL) // Shortens a URL
    http.HandleFunc("/", handleRedirect)          // Redirects short URL to original URL
}
