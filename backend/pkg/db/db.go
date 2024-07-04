
package db

import (
    "database/sql"
    "log"
    

    _ "github.com/mattn/go-sqlite3" // SQLite driver
)

var db *sql.DB // Global variable to hold the database connection

// InitDB initializes the database connection and creates the table if not exists
func InitDB() {
    var err error
    db, err = sql.Open("sqlite3", "./urls.db")
    if err != nil {
        log.Fatal(err)
    }

    // Ensure the database connection is available
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to SQLite database")

    // Create the table if it does not exist
    err = createURLsTable()
    if err != nil {
        log.Fatal(err)
    }
}

// createURLsTable creates the URLs table if it doesn't exist
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

// InsertURL inserts a new URL pair into the database
func InsertURL(shortURL, originalURL string) error {
    _, err := db.Exec("INSERT INTO urls (shortURL, originalURL) VALUES (?, ?)", shortURL, originalURL)
    if err != nil {
        return err
    }
    return nil
}

// GetOriginalURL retrieves the original URL from the database based on the short URL
func GetOriginalURL(shortURL string) (string, error) {
    var originalURL string
    err := db.QueryRow("SELECT originalURL FROM urls WHERE shortURL = ?", shortURL).Scan(&originalURL)
    if err != nil {
        return "", err
    }
    return originalURL, nil
}

// CloseDB closes the database connection
func CloseDB() {
    err := db.Close()
    if err != nil {
        log.Fatal(err)
    }
}
