package main

import (
	"io"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	io.WriteString(w, "Serach: +"+v)
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}
