package main

import (
	"fmt"
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Another dummy log: POST handler single file change test")
	// [PR DEMO 9] Simulate change in POST endpoint for notification test
	fmt.Println("Dummy log: POST handler called for Jenkins diff test")
	msg := getSharedMessage()
	fmt.Fprintf(w, "POST request received. Shared message: %s", msg)
}

func getSharedMessage() string {
	// Shared function between GET and POST handlers
	// [PR DEMO 5] New comment for Jenkins notification test
	return "This is a shared message between GET and POST handlers."
}
