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

func main() {
	var j jugad
	err := http.ListenAndServe(":8080", j)
	if err != nil {
		log.Fatalln(err)
	}
}
