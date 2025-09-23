package main

import (
	"fmt"
	"net/http"
)

// GetHandler depends on PostHandler for some shared logic
func GetHandler(w http.ResponseWriter, r *http.Request) {
	msg := getSharedMessage()
	fmt.Fprintf(w, "GET request received. Shared message: %s", msg)
}

// getSharedMessage is defined in post_handler.go
