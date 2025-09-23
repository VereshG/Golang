package main

import (
	"fmt"
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	msg := getSharedMessage()
	fmt.Fprintf(w, "POST request received. Shared message: %s", msg)
}

func getSharedMessage() string {
	return "This is a shared message between GET and POST handlers."
}
