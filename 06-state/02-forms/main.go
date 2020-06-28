package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	// fmt.Println(req.URL.Path)

	query := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tmpl.Execute(w, query)
	if err != nil {
		log.Println(err)
	}
}
