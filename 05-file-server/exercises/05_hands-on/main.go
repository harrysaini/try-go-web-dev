package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/pics/", http.FileServer(http.Dir("./public")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(res http.ResponseWriter, req *http.Request) {
	tmpl.Execute(res, nil)
}
