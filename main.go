package main

import (
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Retrieve and set default response body
	responseBody := os.Getenv("HTTP_RESPONSE_BODY")
	if responseBody == "" {
		responseBody = "ok" // Default response body if not set
	}

	// Retrieve and set default response code
	responseCode, err := strconv.Atoi(os.Getenv("HTTP_RESPONSE_CODE"))
	if err != nil {
		responseCode = 200 // Default to 200 OK if the env var is not set or invalid
	}

	// Define the handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(responseCode)
		w.Write([]byte(responseBody))
	})

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}
