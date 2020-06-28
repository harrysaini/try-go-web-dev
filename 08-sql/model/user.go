package model

// User type
type User struct {
	First    string
	Last     string
	Username string
	Password string
}

// Session type
type Session struct {
	ID       string
	Username string
}
