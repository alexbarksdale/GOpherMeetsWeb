package main

import (
	"io"
	"net/http"
)

type gopher int

func (g gopher) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/gopher":
		io.WriteString(res, "gopher")
	case "/superGopher":
		io.WriteString(res, "superGopher")
	}
}

func main() {
	var g gopher
	http.ListenAndServe(":8080", g)
}
