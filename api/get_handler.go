package main

import (
	"fmt"
	"net/http"
)

// this API is owned by member core team
func GetHandler(w http.ResponseWriter, r *http.Request) {

	// [PR DEMO 21] Simulate change in GET endpoint for notification test
	fmt.Println("Dummy log: GET handler called for Jenkins diff test")
	msg := getSharedMessage()
	fmt.Fprintf(w, "GET request received. Shared message: %s", msg)
	// [PR DEMO 21] Simulate change in GET endpoint for notification test
}
