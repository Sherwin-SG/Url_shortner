# URL Shortener Project
A simple URL shortener application built with Go and React.

**Author:** Sherwin Gonsalves  
**Date:** July 9, 2024

## Table of Contents
1. [Introduction](#introduction)
2. [Features](#features)
3. [Technologies Used](#technologies-used)
4. [Installation Instructions](#installation-instructions)
5. [Usage Instructions](#usage-instructions)
6. [Database Schema](#database-schema)
7. [Code Documentation](#code-documentation)
8. [Citing Sources](#citing-sources)

## Introduction
The URL Shortener Project is a web application that allows users to shorten long URLs for easier sharing and management. This project demonstrates the use of Go for the backend and React for the frontend.

## Features
- Shorten long URLs
- Redirect to original URLs using the shortened version
- View all shortened URLs
- Delete shortened URLs

## Technologies Used
- Go
- React
- SQLite
- HTML/CSS

## Installation Instructions

### Prerequisites
- Go (version 1.16 or later)
- Node.js (version 14 or later)
- npm (version 6 or later)

## Usage Instructions
1. Open your web browser and navigate to `http://localhost:3000`.
2. Use the interface to input long URLs and get shortened versions.
3. View all shortened URLs in the sliding panel.
4. Delete URLs if no longer needed.

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
