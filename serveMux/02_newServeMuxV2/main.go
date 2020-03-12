package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Gopher")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Super Goph")
}

func main() {

	// /gopher/ - can catch things after because of the '/' after
	http.HandleFunc("/gopher/", d)
	// /superGopher - cannot catch things after (exact match)
	http.HandleFunc("/superGopher", c)
	http.ListenAndServe(":8080", nil) // Uses DefaultServeMux
}
