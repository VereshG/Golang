package main

import (
	"fmt"
	"net/http"
)

// GetHandler depends on PostHandler for some shared logic
// [PR DEMO] This comment added to test Jenkins and PR notification
// [PR DEMO 2] Second demo comment for another commit
func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Dummy log: GET handler called for Jenkins diff test")
	msg := getSharedMessage()
	fmt.Fprintf(w, "GET request received. Shared message: %s", msg)
}

// getSharedMessage is defined in post_handler.go
