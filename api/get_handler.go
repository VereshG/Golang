package main

// get_handler.go
// Handles GET requests for the Go API project.
// This file contains the logic for the GET endpoint and is monitored by Jenkins for changes.
// If this file is changed and merged to the release branch, Jenkins will notify the funds team via Slack.

import (
	"fmt"
	"net/http"
)

// this API is owned by core team
func GetHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Dummy log: GET handler called for Jenkins diff test")
	msg := getSharedMessage()
	fmt.Fprintf(w, "GET request received. Shared message: %s", msg)
}

// [PR DEMO 9] Simulate change in POST endpoint for notification test
