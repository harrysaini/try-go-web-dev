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

	http.HandleFunc("/cat", c)
	http.Handle("/dog", http.HandlerFunc(d))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
