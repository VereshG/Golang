package main

// post_handler.go
// Handles POST requests for the Go API project.
// This file contains the logic for the POST endpoint and is monitored by Jenkins for changes.
// If this file is changed and merged to the release branch, Jenkins will notify the core team via Slack.

import (
	"fmt"
	"net/http"
)

// this API is owned by member funds team
func PostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Another dummy log: POST handler single file change test")

	fmt.Println("Dummy log: POST handler called for Jenkins diff test")
	msg := getSharedMessage()
	fmt.Fprintf(w, "POST request received. Shared message: %s", msg)
	fmt.Fprintf(w, "POST request received. Shared message: %s", msg)
	// [PR DEMO 21] Simulate change in POST endpoint for notification tests
}

func getSharedMessage() string {

	return "This is a shared message between GET and POST handlers."
}
