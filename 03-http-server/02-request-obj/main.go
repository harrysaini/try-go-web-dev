package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type jugad int

func (j jugad) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		log.Fatalln(err)
	}

	// fmt.Println(req.PostForm)

	data := struct {
		Method        string
		Submissions   url.Values
		Header        http.Header
		URL           *url.URL
		ContentLength int64
		Host          string
	}{
		req.Method,
		req.Form,
		req.Header,
		req.URL,
		req.ContentLength,
		req.Host,
	}

	res.Header().Set("my-header", "my-value")
	res.WriteHeader(http.StatusBadRequest)
	tmpl.ExecuteTemplate(res, "tmpl.gohtml", data)
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {
	var j jugad
	err := http.ListenAndServe(":8080", j)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Server listening")
}
