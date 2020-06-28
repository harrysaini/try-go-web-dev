package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/about/", about)
	http.HandleFunc("/contact/", contact)
	http.HandleFunc("/apply", applyMain)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func applyMain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		apply(w, req)
	case "POST":
		applyProcess(w, req)
	}
}

func apply(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

func applyProcess(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
	HandleError(w, err)
}

// HandleError is used to handle error
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
