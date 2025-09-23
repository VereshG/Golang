package main

import (
	"fmt"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	// getSharedMessage is defined in post_handler.go
	// getSharedMessage is defined in post_handler.go
	// getSharedMessage is defined in post_handler.go
	// [PR DEMO 5] New comment for Jenkins notification test
	// [PR DEMO 8] Simulate change in GET endpoint for notification test
	fmt.Println("Dummy log: GET handler called for Jenkins diff test")
	msg := getSharedMessage()
	fmt.Fprintf(w, "GET request received. Shared message: %s", msg)
}

// getSharedMessage is defined in post_handler.go
