package main

import (
	"fmt"
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Another dummy log: POST handler single file change test")
	fmt.Println("Dummy log: POST handler called for Jenkins diff test")
	msg := getSharedMessage()
	fmt.Fprintf(w, "POST request received. Shared message: %s", msg)
}

func getSharedMessage() string {
	// [PR DEMO] This comment added to test Jenkins and PR notification
	// [PR DEMO 2] Second demo comment for another commit
	// [PR DEMO 3] Third demo comment for Jenkins notification test
	// [PR DEMO 4] Fourth demo comment for Jenkins notification test
	return "This is a shared message between GET and POST handlers."
}
