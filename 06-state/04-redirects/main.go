package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/redirectme", redirectme)
	http.HandleFunc("/redirected", redirected)
	http.HandleFunc("/foo", redirectme2)
	http.HandleFunc("/bar", redirected)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Index")
}

func redirectme(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Location", "/redirected")
	w.WriteHeader(http.StatusMovedPermanently)
}

func redirectme2(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/bar", http.StatusSeeOther)
}

func redirected(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Redirected")
}
