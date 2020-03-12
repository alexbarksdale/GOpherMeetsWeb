package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func dog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src ="toby.jpg" />`)
}

func dogPic(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		HandleError(res, err)
	}
	defer f.Close()

	fi, err := f.Stat() // Stat() - stats of the file
	if err != nil {
		HandleError(res, err)
	}

	http.ServeContent(res, req, f.Name(), fi.ModTime(), f)

}

func HandleError(res http.ResponseWriter, err error) {
	http.Error(res, "file not found", 404)
	log.Fatalln(err)
	return
}

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}
