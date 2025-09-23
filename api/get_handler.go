package main

import (
	"fmt"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	// getSharedMessage is defined in post_handler.go
	// getSharedMessage is defined in post_handler.go
	// getSharedMessage is defined in post_handler.go
	fmt.Println("Dummy log: GET handler called for Jenkins diff test")
	msg := getSharedMessage()
	fmt.Fprintf(w, "GET request received. Shared message: %s", msg)
}

// getSharedMessage is defined in post_handler.go
