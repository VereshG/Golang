package main

import (
	"fmt"
	"net/http"
)

// GetHandler depends on PostHandler for some shared logic
// [PR DEMO] This comment added to test Jenkins and PR notification
// [PR DEMO 2] Second demo comment for another commit
// [PR DEMO 3] Third demo comment for Jenkins notification test
// [PR DEMO 4] Fourth demo comment for Jenkins notification test
func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Dummy log: GET handler called for Jenkins diff test")
	msg := getSharedMessage()
	fmt.Fprintf(w, "GET request received. Shared message: %s", msg)
}

// getSharedMessage is defined in post_handler.go
