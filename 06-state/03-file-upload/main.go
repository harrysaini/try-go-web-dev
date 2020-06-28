package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

// Data to pass to template
type Data struct {
	File string
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	// fmt.Println(req.URL.Path)

	var data Data

	if req.Method == http.MethodPost {

		file, fileHeader, err := req.FormFile("q")

		willUpload := req.FormValue("upload") == "on"

		fmt.Println(willUpload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// fmt.Println(file)
		// fmt.Println(fileHeader)

		fileContent, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if willUpload {
			dest, err := os.Create(filepath.Join("./uploads", fileHeader.Filename))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			n, err := dest.Write(fileContent)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			fmt.Println(n)
		}

		data = Data{
			string(fileContent),
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}
