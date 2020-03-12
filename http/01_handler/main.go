package main

import (
	"fmt"
	"net/http"
)

type gopher int

// Gopher is now also type Handler
func (g gopher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Anything in this func")
}

func main() {
	var d gopher
	http.ListenAndServe(":8080", d) // Takes in a handler (d - gopher)

}
