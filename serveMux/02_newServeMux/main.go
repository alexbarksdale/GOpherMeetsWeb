package main

import (
	"io"
	"net/http"
)

type gopher int

func (g gopher) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Gopher")
}

type superGopher int

func (s superGopher) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Super Goph")
}

func main() {
	var g gopher
	var s superGopher

	mux := http.NewServeMux()
	// /gopher/ - can catch things after because of the '/' after
	mux.Handle("/gopher/", g)
	// /superGopher - cannot catch things after (exact match)
	mux.Handle("/superGopher", s)
	http.ListenAndServe(":8080", mux)
}
