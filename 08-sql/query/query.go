package query

import (
	"database/sql"
	"log"

	"github.com/harrysaini/try-go-web-dev/08-sql/model"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:olx123@tcp(127.0.0.1:3306)/test-db-go?charset=utf8")

	if err != nil {
		log.Fatalln(err)
	}
}

// GetUser get user from database
func GetUser(username string) (*model.User, error) {
	stmt, err := db.Prepare("SELECT `username`, `first`, `last`, `password` from `users` where `username`=?")
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(username)

	var password, first, last string

	err = row.Scan(&username, &first, &last, &password)
	if err != nil {
		return nil, err
	}

	user := model.User{
		First:    first,
		Last:     last,
		Password: password,
		Username: username,
	}

	return &user, nil

}

// GetSession get session from database
func GetSession(sessionID string) (*model.Session, error) {
	stmt, err := db.Prepare("SELECT * from `sessions` where `id`=?")
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(sessionID)

	var username string

	err = row.Scan(&sessionID, &username)
	if err != nil {
		return nil, err
	}

	session := model.Session{
		ID:       sessionID,
		Username: username,
	}

	return &session, nil

}

// SaveUser save user into database
func SaveUser(user *model.User) error {
	stmt, err := db.Prepare("INSERT INTO `users` (`username`,`first`,`last`,`password`) VALUES (?,?,?,?)")

	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Username, user.First, user.Last, user.Password)
	if err != nil {
		return err
	}

	num, err := result.RowsAffected()
	if err != nil || num == 0 {
		return err
	}

	return nil

}

// SaveSession save session into database
func SaveSession(session *model.Session) error {
	stmt, err := db.Prepare("INSERT INTO `sessions` (`id`, `username`) VALUES (?, ?)")

	if err != nil {
		return err
	}

	result, err := stmt.Exec(session.ID, session.Username)
	if err != nil {
		return err
	}

	num, err := result.RowsAffected()
	if err != nil || num == 0 {
		return err
	}

	return nil

}
