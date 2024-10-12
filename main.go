package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Go web server!")
}

func hahaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Haha! You found the funny route!")
}

func main() {
	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/haha", hahaHandler)

	// Start the server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
