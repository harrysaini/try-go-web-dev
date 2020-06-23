package main

import (
	"fmt"
	"log"
	"net/http"
)

type jugad int

func (j jugad) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello world")
}

type cat int

func (j cat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello cat")
}

func main() {
	var j jugad
	var c cat

	mux := http.NewServeMux()

	mux.Handle("/cat", c)
	mux.Handle("/dog", j)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalln(err)
	}
}
