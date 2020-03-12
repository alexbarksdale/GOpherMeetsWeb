package main

import (
	"log"
	"net/http"
)

func main() {
	// Handles error if any
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
