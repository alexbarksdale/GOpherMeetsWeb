package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type gopher int

// Gopher is now also type Handler
func (g gopher) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm() // Must ParseForm() before getting the values
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		URL         *url.URL
		Submissions url.Values
	}{
		req.Method,
		req.URL,
		req.Form,
	}

	tpl.ExecuteTemplate(res, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d gopher
	http.ListenAndServe(":8080", d) // Gets handled by ServeHTTP

}
