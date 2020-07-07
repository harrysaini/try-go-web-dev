package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

// Helper

func getCookie(req *http.Request) *http.Cookie {
	cookie, err := req.Cookie("session")
	if err != nil {
		uuid, _ := uuid.NewRandom()
		cookie = &http.Cookie{
			Name:  "session",
			Value: uuid.String(),
		}
	}

	return cookie
}

func appendValueInCookie(cookie *http.Cookie, fileName string) {
	if strings.Contains(cookie.Value, fileName) {
		return
	}

	cookie.Value += ("|" + fileName)
}

// Server

func main() {
	http.HandleFunc("/", index)
	http.Handle("/uploads/", http.StripPrefix("/uploads", http.FileServer(http.Dir("./uploads/"))))
	http.Handle("favicon.ico", http.NotFoundHandler())

	log.Fatalln(http.ListenAndServe(":8090", nil))

}

func index(w http.ResponseWriter, req *http.Request) {
	cookie := getCookie(req)

	if req.Method == http.MethodPost {
		uploadFile, uploadFH, err := req.FormFile("file")

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		hash := sha1.New()

		io.Copy(hash, uploadFile)

		hashVal := fmt.Sprintf("%x", hash.Sum(nil))

		fileName := hashVal + "." + strings.Split(uploadFH.Filename, ".")[1]

		file, err := os.Create(filepath.Join("./uploads/", fileName))
		if err != nil {
			log.Println(err)
		}

		uploadFile.Seek(0, 0)

		io.Copy(file, uploadFile)

		appendValueInCookie(cookie, fileName)

	}

	http.SetCookie(w, cookie)

	xs := strings.Split(cookie.Value, "|")

	tmpl.Execute(w, xs)

}
