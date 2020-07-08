package lib

import (
	"database/sql"
	"log"

	"github.com/harrysaini/try-go-web-dev/012-posgres/dblink"

	// postgres driver
	_ "github.com/lib/pq"
)

// DB reference to sql db connection
var DB *sql.DB

// DataStore Provide methods to interract with database
type DataStore struct {
	BooksDataStore *dblink.BooksSQLDataStore
}

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

// InitDataStores data link with sql
func InitDataStores() *DataStore {
	booksDataStore := &dblink.BooksSQLDataStore{
		DB: DB,
	}

	return &DataStore{
		booksDataStore,
	}
}
