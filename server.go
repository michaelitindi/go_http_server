package main

import (
	"fmt"
	"net/http"
)
func StartServer() {
    // Set up route handlers
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/about", aboutHandler)
    http.HandleFunc("/users", usersHandler)

    // Start the server
    fmt.Println("Server starting on http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}

// Handler for the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintf(w, `
        <html>
            <body>
                <h1>Welcome to Go HTTP Server</h1>
                <p>This is the home page.</p>
                <ul>
                    <li><a href="/">Home</a></li>
                    <li><a href="/about">About</a></li>
                    <li><a href="/users">Users</a></li>
                </ul>
            </body>
        </html>
    `)
}

// Handler for the about page
func aboutHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintf(w, `
        <html>
            <body>
                <h1>About Page</h1>
                <p>This is a simple HTTP server built with Go.</p>
                <a href="/">Back to Home</a>
            </body>
        </html>
    `)
}

// Handler for users list
func usersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintf(w, `
        <html>
            <body>
                <h1>Users</h1>
                <ul>
                    <li>Alice</li>
                    <li>Bob</li>
                    <li>Charlie</li>
                </ul>
                <a href="/">Back to Home</a>
            </body>
        </html>
    `)
}