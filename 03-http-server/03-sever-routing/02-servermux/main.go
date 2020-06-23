package main

import (
	"fmt"
	"log"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello cat")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/cat", c)
	mux.HandleFunc("/dog", d)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalln(err)
	}
}
