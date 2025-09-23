package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("Received request from %s\n", r.RemoteAddr)
	num := rand.Intn(1000)
	fmt.Fprintf(w, "Hello! Your random number is: %d", num)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080...")
}
