// package server
package main

import (
	"fmt"
	"net/http"
)

// handler function for root ("/")
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page")
}

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
