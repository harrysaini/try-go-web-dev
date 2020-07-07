package lib

import (
	"database/sql"
	"log"

	// postgres driver
	_ "github.com/lib/pq"
)

// DB reference to sql db connection
var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://harish:olx123@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

}
