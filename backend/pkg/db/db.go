package db

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3" 
)

var db *sql.DB 

type URL struct {
    Original  string `json:"originalUrl"`
    Shortened string `json:"shortenedUrl"`
}

func InitDB() {
    var err error
    db, err = sql.Open("sqlite3", "./urls.db")
    if err != nil {
        log.Fatal(err)
    }

   
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to SQLite database")

    err = createURLsTable()
    if err != nil {
        log.Fatal(err)
    }
}

func createURLsTable() error {
    createTableSQL := `
        CREATE TABLE IF NOT EXISTS urls (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            shortURL TEXT UNIQUE NOT NULL,
            originalURL TEXT NOT NULL
        );
    `

    _, err := db.Exec(createTableSQL)
    if err != nil {
        return err
    }

    log.Println("Created URLs table if not exists")
    return nil
}


func InsertURL(shortURL, originalURL string) error {
    _, err := db.Exec("INSERT INTO urls (shortURL, originalURL) VALUES (?, ?)", shortURL, originalURL)
    if err != nil {
        return err
    }
    return nil
}


func GetOriginalURL(shortURL string) (string, error) {
    var originalURL string
    err := db.QueryRow("SELECT originalURL FROM urls WHERE shortURL = ?", shortURL).Scan(&originalURL)
    if err != nil {
        return "", err
    }
    return originalURL, nil
}

func GetAllURLs() ([]URL, error) {
    rows, err := db.Query("SELECT shortURL, originalURL FROM urls")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var urls []URL
    for rows.Next() {
        var url URL
        err := rows.Scan(&url.Shortened, &url.Original)
        if err != nil {
            return nil, err
        }
        urls = append(urls, url)
    }

    return urls, nil
}


func CloseDB() {
    err := db.Close()
    if err != nil {
        log.Fatal(err)
    }
}
