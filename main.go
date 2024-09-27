package main

import (
    "log"
    "net/http"
	"github.com/JawherKl/short-basic-url"
)

func main() {
    // Register routes
    routes()

    // Start server on port 8080
    log.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}
