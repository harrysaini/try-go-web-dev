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
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<html> Index <br/> <a href=\"/set\"> SET </a></html>")
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "Cookie-custom",
		Value: "Hello",
	})
	fmt.Fprintln(w, "<html>COOKIE WRITTEN - CHECK YOUR BROWSER <a href=\"/read\"> Read </a> </html>")

}

func read(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("Cookie-custom")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}

	tmpl.Execute(w, cookie)
}

func expire(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("Cookie-custom")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}

	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	http.Redirect(w, req, "/", http.StatusSeeOther)

}
