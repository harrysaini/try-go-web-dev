package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("something.gohtml"))
}

func d(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, req.URL.Path)
	tmpl.Execute(res, "dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, req.URL.Path)
	tmpl.Execute(res, "Index")
}

func me(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, req.URL.Path)
	tmpl.Execute(res, "Harish")
}

func main() {

	http.Handle("/dog", http.HandlerFunc(d))
	http.HandleFunc("/me", me)
	http.HandleFunc("/", c)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
