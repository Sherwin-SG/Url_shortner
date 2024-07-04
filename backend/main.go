package main

import (
    "encoding/json"
    "log"
    "math/rand"
    "net/http"
    "fmt"

    "github.com/rs/cors"
    "react-go-project/backend/pkg/db"
)

func main() {
    db.InitDB()
    defer db.CloseDB()

    mux := http.NewServeMux()

    // Handle API endpoints
    mux.HandleFunc("/api/shorten", shortenURLHandler)
    
    // Handle redirect for shortened URLs
    mux.HandleFunc("/", redirectHandler)
    
    // Default handler for other routes
    mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to URL Shortener API")
    })

    handler := cors.Default().Handler(mux)

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", handler))
}

// URL struct for encoding and decoding JSON
type URL struct {
    Original  string `json:"originalUrl"`
    Shortened string `json:"shortenedUrl"`
}

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
    var url URL
    err := json.NewDecoder(r.Body).Decode(&url)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Generate shortURL
    shortURL := generateShortURL()
    log.Println("Generated short URL:", shortURL)

    // Insert into database (assuming db.InsertURL function exists)
    err = db.InsertURL(shortURL, url.Original)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Prepare response with shortened URL
    url.Shortened = "http://localhost:8080/" + shortURL
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(url)
}



func redirectHandler(w http.ResponseWriter, r *http.Request) {
    shortURL := r.URL.Path[1:] // Extract the short URL from the request path
    log.Println("Received short URL:", shortURL)

    originalURL, err := db.GetOriginalURL(shortURL)
    if err != nil {
        log.Println("Error retrieving original URL:", err)
        http.Error(w, "URL not found", http.StatusNotFound)
        return
    }

    // Redirect to the original URL
    http.Redirect(w, r, originalURL, http.StatusFound)
}

func generateShortURL() string {
    const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    b := make([]byte, 6)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}
