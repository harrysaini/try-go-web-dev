package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var tmpl *template.Template

func init() {
	path, _ := filepath.Abs("./templates/index.gohtml")
	fmt.Println(path)
	tmpl = template.Must(template.ParseFiles(path))
}

func main() {
	path, _ := filepath.Abs("./public")
	fmt.Println(path)
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir(path))))

	log.Fatal(http.ListenAndServe(":80", nil))
}

func index(res http.ResponseWriter, req *http.Request) {
	tmpl.Execute(res, nil)
}
