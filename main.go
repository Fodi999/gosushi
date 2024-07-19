package main

import (
    "fmt"
    "net/http"
)

// Handler for the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

// Handler for the API endpoint
func apiHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, `{"message": "Hello, world!"}`)
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/api/message", apiHandler)

    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
