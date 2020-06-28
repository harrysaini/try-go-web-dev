package main

import (
	"fmt"
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
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "Cookie-custom",
		Value: "Hello",
	})
	fmt.Fprintln(w, "Index")
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")

}

func read(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("Cookie-custom")
	if err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(w, cookie)
}
