package main

import (
	"fmt"
	"net/http"
)

type gopher int

func (g gopher) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// Inspect element to header to see below
	res.Header().Set("Cool-Key", "This is from cool-key")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(res, "<h1>Things can go in here</h1>")
}

func main() {
	var d gopher
	http.ListenAndServe(":8080", d) // passed in the handler (ServeHTTP)
}
