package main

import (
	"html/template"
	"log"
	"net/http"
)

type gopher int

// Gopher is now also type Handler
func (g gopher) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm() // Must ParseForm() before getting the values
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d gopher
	http.ListenAndServe(":8080", d) // Gets handled by ServeHTTP

}
