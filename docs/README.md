# URL Shortener Project
A simple URL shortener application built with Go and React.

**Author:** Sherwin Gonsalves  
**Date:** July 9, 2024

## Table of Contents
1. [Introduction](#introduction)
2. [Installation Instructions](#installation-instructions)
3. [Usage Instructions](#usage-instructions)
5. [Database Schema](#Database-Schema)
6. [Code Documentation](#Code-Documentation)
7. [Citing Sources](#citing-sources)


## Introduction
The URL Shortener Project is a web application that allows users to shorten long URLs for easier sharing and management. This project demonstrates the use of Go for the backend and React for the frontend.

### Features
- Shorten long URLs
- Redirect to original URLs using the shortened version
- View all shortened URLs
- Delete shortened URLs

### Technologies Used
- Go
- React
- SQLite
- HTML/CSS

## Installation Instructions

### Prerequisites
- Go (version 1.16 or later)
- Node.js (version 14 or later)
- npm (version 6 or later)


**Database Schema:**
```markdown
## Database Schema

### URLs Table
- **id:** INTEGER PRIMARY KEY AUTOINCREMENT
- **shortURL:** TEXT UNIQUE NOT NULL
- **originalURL:** TEXT NOT NULL


## Code Documentation

### Backend

#### main.go
The main entry point for the backend server.

#### db/db.go
Handles database connections and operations.

#### Handlers
- **shortenURLHandler:** Handles URL shortening requests.
- **redirectHandler:** Handles redirection to the original URL.
- **getAllURLsHandler:** Retrieves all shortened URLs.
- **deleteURLHandler:** Deletes a shortened URL.

### Frontend

#### urlshortener.js
Main component for the URL shortener interface.

#### slidingpanel.js
Component for displaying and managing all shortened URLs.

#### copybutton.js
Component for copying the shortened URL to the clipboard.


## Citing Sources

### Libraries and Tools
- **Go:** https://golang.org/
- **React:** https://reactjs.org/
- **SQLite:** https://www.sqlite.org/
- **cors:** https://github.com/rs/cors
- **go-sqlite3:** https://github.com/mattn/go-sqlite3

### Tutorials and Articles
- [React Documentation](https://reactjs.org/docs/getting-started.html)
