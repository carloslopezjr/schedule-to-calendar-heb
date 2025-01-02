package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Channel to store authorization codes (optional for testing/debugging)
	authCodeChan := make(chan string)

	// HTTP handler to capture the authorization code
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code != "" {
			fmt.Fprintln(w, "Authorization successful! You can close this window.")
			fmt.Printf("Received authorization code: %s\n", code)
			authCodeChan <- code // Send code to channel for debugging or external use
		} else {
			fmt.Fprintln(w, "Authorization failed or no code received!")
		}
	})

	// Start the server on port 80
	fmt.Println("Authorization server started on http://localhost:80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
