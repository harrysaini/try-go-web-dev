package models

import "errors"

var (
	BooksFieldsUdateAllowed = []string{"name", "author", "price"}
)

// Book ds
type Book struct {
	Name   string `json:"name"`
	Isbn   string `json:"isbn"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

// Validate book data
func (book *Book) Validate() (bool, error) {
	if book.Isbn == "" {
		return false, errors.New("Isbn should be present")
	}
	if book.Name == "" {
		return false, errors.New("Name should be present")
	}
	if book.Author == "" {
		return false, errors.New("Author should be present")
	}
	if book.Price == 0 {
		return false, errors.New("Price should be present")
	}

	return true, nil
}

// BooksDataStore represent all operations needs
type BooksDataStore interface {
	GetAllBooks() (books *[]Book, err error)
	GetBook(isbn string) (book *Book, err error)
	InsertBook(book *Book) (ok bool, err error)
	UpdateBook(isbn string, updates map[string]interface{}) (ok bool, err error)
	DeleteBook(isbn string) (ok bool, err error)
}
