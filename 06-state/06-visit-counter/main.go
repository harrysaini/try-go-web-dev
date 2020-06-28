package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	var count int64
	cookie, err := req.Cookie("count")
	if err != nil {
		fmt.Println(err)
		count = 1
	} else {
		count, err = strconv.ParseInt(cookie.Value, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(count)
		count = count + 1
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "count",
		Value: strconv.FormatInt(count, 10),
	})
	tmpl.Execute(w, count)

}
