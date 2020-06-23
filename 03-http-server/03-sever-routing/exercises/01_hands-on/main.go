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
	fmt.Fprintln(res, "Hello Index")
}

func me(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Harish")
}

func main() {

	http.HandleFunc("/", c)
	http.Handle("/dog/", http.HandlerFunc(d))
	http.HandleFunc("/me", me)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
