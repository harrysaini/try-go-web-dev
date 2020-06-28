package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var tmpl *template.Template

// User type
type User struct {
	First    string
	Last     string
	Username string
	Password []byte
}

var dbUsers = map[string]User{}
var dbSessions = map[string]string{}

func init() {
	tmpl = template.Must(template.ParseGlob("./templates/*"))
}

// Helpers code

func getUserForRequest(req *http.Request) (user *User, ok bool) {
	cookie, err := req.Cookie("session")
	if err != nil {
		return nil, false
	}

	sessionID := cookie.Value

	userName, ok := dbSessions[sessionID]
	if !ok {
		return nil, false
	}

	userStruct, ok := dbUsers[userName]
	return &userStruct, ok

}

func createSession() (*http.Cookie, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	var cookie = http.Cookie{
		Name:  "session",
		Value: uuid.String(),
	}

	return &cookie, err
}

// Server code

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/read", read)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	user, _ := getUserForRequest(req)
	tmpl.ExecuteTemplate(w, "index.gohtml", user)

}

func signup(w http.ResponseWriter, req *http.Request) {
	_, ok := getUserForRequest(req)
	if ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")
		first := req.FormValue("first")
		last := req.FormValue("last")

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		user := User{
			Username: username,
			Password: hash,
			First:    first,
			Last:     last,
		}

		dbUsers[username] = user

		cookie, err := createSession()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		dbSessions[cookie.Value] = username

		http.SetCookie(w, cookie)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tmpl.ExecuteTemplate(w, "signup.gohtml", nil)

}

func login(w http.ResponseWriter, req *http.Request) {
	_, ok := getUserForRequest(req)
	if ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")

		user, ok := dbUsers[username]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		cookie, err := createSession()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		dbSessions[cookie.Value] = username

		http.SetCookie(w, cookie)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tmpl.ExecuteTemplate(w, "login.gohtml", nil)

}

func read(w http.ResponseWriter, req *http.Request) {
	user, ok := getUserForRequest(req)
	if !ok {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	tmpl.Execute(w, *user)
}

func logout(w http.ResponseWriter, req *http.Request) {
	_, ok := getUserForRequest(req)
	if !ok {
		http.Redirect(w, req, "/login", http.StatusBadRequest)
		return
	}

	cookie, _ := req.Cookie("session")
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	http.Redirect(w, req, "/", http.StatusSeeOther)

}
