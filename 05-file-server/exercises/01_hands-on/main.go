package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/code", code)
	http.HandleFunc("/code.jpg", getImage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello world")
}

func code(res http.ResponseWriter, req *http.Request) {
	tmpl.Execute(res, nil)
}

func getImage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./assets/code.png")
}
