package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	query := req.FormValue("q")
	io.WriteString(w, "Do your search: "+query)
}
