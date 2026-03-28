package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/submit", handleSubmission)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func handleSubmission(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Submission received!")
}
