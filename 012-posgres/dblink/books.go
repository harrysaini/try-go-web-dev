package dblink

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/harrysaini/try-go-web-dev/012-posgres/models"
)

// BooksSQLDataStore provides linking with db
type BooksSQLDataStore struct {
	DB *sql.DB
}

// GetAllBooks get all books
func (ds *BooksSQLDataStore) GetAllBooks() (*[]models.Book, error) {
	rows, err := ds.DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]models.Book, 0)

	for rows.Next() {
		bk := models.Book{}
		err := rows.Scan(&bk.Isbn, &bk.Name, &bk.Author, &bk.Price) // order matters
		if err != nil {
			return nil, err
		}

		books = append(books, bk)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &books, nil
}

// InsertBook add new book
func (ds *BooksSQLDataStore) InsertBook(book *models.Book) (bool, error) {
	result, err := ds.DB.Exec(
		"INSERT INTO books(id, name, author, price) VALUES ($1, $2, $3, $4)",
		book.Isbn, book.Name, book.Author, book.Price)
	if err != nil {
		return false, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, errors.New("Unable to insert data")
	}

	return true, nil
}

// GetBook get book by isbn
func (ds *BooksSQLDataStore) GetBook(isbn string) (*models.Book, error) {
	var book models.Book

	err := ds.DB.QueryRow("SELECT * FROM books WHERE id=$1", isbn).Scan(&book.Isbn, &book.Name, &book.Author, &book.Price)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, errors.New("Not found")
		default:
			return nil, err

		}
	}

	return &book, nil
}

// UpdateBook update book by isbn
func (ds *BooksSQLDataStore) UpdateBook(isbn string, updates map[string]interface{}) (bool, error) {
	var updateQueries []string
	values := []interface{}{isbn}
	counter := 2

	for key, val := range updates {
		updateQueries = append(updateQueries, fmt.Sprintf("%s = $%d", key, counter))
		values = append(values, val)
		counter++
	}

	query := fmt.Sprintf("Update books SET %s WHERE id=$1", strings.Join(updateQueries, ","))

	log.Println(query)

	result, err := ds.DB.Exec(query, values...)
	if err != nil {
		return false, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, errors.New("Unable to insert data")
	}

	return true, nil
}

// DeleteBook get book by isbn
func (ds *BooksSQLDataStore) DeleteBook(isbn string) (bool, error) {
	res, err := ds.DB.Exec("DELETE FROM books WHERE id=$1", isbn)
	if err != nil {
		return false, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, errors.New("Unable to delete data")
	}

	return true, nil

}
