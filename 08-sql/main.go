package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/harrysaini/try-go-web-dev/08-sql/model"
	"github.com/harrysaini/try-go-web-dev/08-sql/query"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("./templates/*"))
}

// Helpers code

func getUserForRequest(req *http.Request) (user *model.User, ok bool) {
	cookie, err := req.Cookie("session")
	if err != nil {
		return nil, false
	}

	sessionID := cookie.Value

	session, err := query.GetSession(sessionID)
	if err != nil {
		return nil, false
	}

	username := session.Username

	fmt.Println("session", session)

	user, err = query.GetUser(username)
	if err != nil {
		return nil, false
	}

	fmt.Println("user", user)
	return user, true

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
	err := tmpl.ExecuteTemplate(w, "index.gohtml", user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

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
			return
		}

		user := model.User{
			Username: username,
			Password: string(hash),
			First:    first,
			Last:     last,
		}

		err = query.SaveUser(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cookie, err := createSession()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session := model.Session{
			ID:       cookie.Value,
			Username: username,
		}

		err = query.SaveSession(&session)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

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

		user, err := query.GetUser(username)
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		fmt.Println(user, password)

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			http.Error(w, "Username and/or password do not match error", http.StatusForbidden)
			return
		}

		cookie, err := createSession()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		session := model.Session{
			ID:       cookie.Value,
			Username: username,
		}

		err = query.SaveSession(&session)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

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
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	cookie, _ := req.Cookie("session")
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	http.Redirect(w, req, "/", http.StatusSeeOther)

}
