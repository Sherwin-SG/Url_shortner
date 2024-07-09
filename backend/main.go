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

   
    mux.HandleFunc("/api/shorten", shortenURLHandler)
    
   
    mux.HandleFunc("/", redirectHandler)
    
   
    mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to URL Shortener API")
    })

    mux.HandleFunc("/api/urls", getAllURLsHandler)

    handler := cors.Default().Handler(mux)

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", handler))
}

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
    var url db.URL
    err := json.NewDecoder(r.Body).Decode(&url)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        log.Printf("Error decoding JSON: %v\n", err)
        return
    }

    shortURL := generateShortURL()
    log.Println("Generated short URL:", shortURL)

    err = db.InsertURL(shortURL, url.Original)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        log.Printf("Error inserting URL into database: %v\n", err)
        return
    }

    url.Shortened = "http://localhost:8080/" + shortURL
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(url)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
    shortURL := r.URL.Path[1:] 
    log.Println("Received short URL:", shortURL)

    originalURL, err := db.GetOriginalURL(shortURL)
    if err != nil {
        log.Println("Error retrieving original URL:", err)
        http.Error(w, "URL not found", http.StatusNotFound)
        return
    }

    
    http.Redirect(w, r, originalURL, http.StatusFound)
}

func getAllURLsHandler(w http.ResponseWriter, r *http.Request) {
    urls, err := db.GetAllURLs()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        log.Printf("Error retrieving URLs from database: %v\n", err)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(urls)
}

func generateShortURL() string {
    const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    b := make([]byte, 6)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}
