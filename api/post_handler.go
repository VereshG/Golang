package main

import (
	"fmt"
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Dummy log: POST handler called for Jenkins diff test")
	msg := getSharedMessage()
	fmt.Fprintf(w, "POST request received. Shared message: %s", msg)
}

func getSharedMessage() string {
	return "This is a shared message between GET and POST handlers."
}
